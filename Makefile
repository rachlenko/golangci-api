.PHONY: test

gen: gen_services gen_models

gen_models:
	go generate ./...

gen_services:
	GO111MODULE=off time go run cmd/genservices/main.go -service=${S}

run_env:
	SERVICES=sqs localstack start --docker

prepare_env:
	awslocal sqs create-queue --queue-name primary
	awslocal sqs list-queues

run_api_reload:
	godotenv gin -i --build cmd/golangci-api run main.go

run_api:
	godotenv go run cmd/golangci-api/main.go

run_worker_reload:
	godotenv gin -i --port 3099 --build cmd/golangci-worker run main.go

run_worker:
	godotenv go run cmd/golangci-worker/main.go

migrate_force_version:
	godotenv -f .env sh -c 'migrate -database $${DATABASE_URL} -path ./migrations force $${V}'

prepare_test_env:
	./test/prepare_env.sh

test_api:
	go test -v -parallel 1 -p 1 ./test/

test_lint:
	golangci-lint run -v -j4

test_lint_dev:
	go run ../golangci-lint/cmd/golangci-lint/main.go run -v

test_api_dev:
	echo "DROP DATABASE api_test;" | docker-compose exec -T pg psql -U postgres
	echo "CREATE DATABASE api_test;" | docker-compose exec -T pg psql -U postgres
	echo "FLUSHALL" | docker-compose exec -T redis redis-cli
	make test_api

test_worker:
	go test -v -parallel 1 -p 1 ./pkg/worker/...

test_worker_dev: test_worker

test: test_api test_worker
test_dev: test_api_dev test_worker_dev test_lint_dev

connect_to_local_db:
	dc exec pg psql -U postgres -d api_prod

build_lambda:
	GOOS=linux go build -o sqsLambdaConsumer ./deployments/awslambda/sqsconsumer/
	zip sqsLambdaConsumer.zip sqsLambdaConsumer

deploy_lambda: build_lambda
	aws s3 cp ./sqsLambdaConsumer.zip s3://golangci-lambda-functions/

worker_test_repo:
	# set env vars PR, REPO
	SLOW_TESTS_ENABLED=1 go test -v ./analyze -run TestAnalyzeRepo

worker_test_repo_fake_github:
	# set env vars PR, REPO
	SLOW_TESTS_ENABLED=1 go test -v ./analyze/processors -count=1 -run TestProcessRepoWithFakeGithub

mod_update:
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

psql:
	docker-compose exec pg psql -U postgres -dapi_prod

build_docker_images:
	cp  ~/deploy/golangci-worker/Dockerfile worker.dockerfile
	docker build -t golangci/golangci-worker -f worker.dockerfile .
	docker push golangci/golangci-worker

	cp  ~/deploy/golangci-api/Dockerfile api.dockerfile
	docker build -t golangci/golangci-api -f api.dockerfile .
	docker push golangci/golangci-api
