// Code generated by genservices. DO NOT EDIT.
package pranalysis

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/golangci/golangci-api/internal/api/apierrors"
	"github.com/golangci/golangci-api/internal/api/endpointutil"
	"github.com/golangci/golangci-api/internal/api/transportutil"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func RegisterHandlers(svc Service, r *mux.Router, regCtx *endpointutil.HandlerRegContext) {

	hGetAnalysisStateByAnalysisGUID := httptransport.NewServer(
		makeGetAnalysisStateByAnalysisGUIDEndpoint(svc, regCtx.Log),
		decodeGetAnalysisStateByAnalysisGUIDRequest,
		encodeGetAnalysisStateByAnalysisGUIDResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreInternalRequestContext(*regCtx)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	r.Methods("GET").Path("/v1/repos/{provider}/{owner}/{name}/analyzes/{analysisguid}/state").Handler(hGetAnalysisStateByAnalysisGUID)

	hGetAnalysisStateByPRNumber := httptransport.NewServer(
		makeGetAnalysisStateByPRNumberEndpoint(svc, regCtx.Log),
		decodeGetAnalysisStateByPRNumberRequest,
		encodeGetAnalysisStateByPRNumberResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAnonymousRequestContext(*regCtx)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	r.Methods("GET").Path("/v1/repos/{provider}/{owner}/{name}/pulls/{pullrequestnumber}").Handler(hGetAnalysisStateByPRNumber)

	hUpdateAnalysisStateByAnalysisGUID := httptransport.NewServer(
		makeUpdateAnalysisStateByAnalysisGUIDEndpoint(svc, regCtx.Log),
		decodeUpdateAnalysisStateByAnalysisGUIDRequest,
		encodeUpdateAnalysisStateByAnalysisGUIDResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreInternalRequestContext(*regCtx)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	r.Methods("PUT").Path("/v1/repos/{provider}/{owner}/{name}/analyzes/{analysisguid}/state").Handler(hUpdateAnalysisStateByAnalysisGUID)

}

func decodeGetAnalysisStateByAnalysisGUIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetAnalysisStateByAnalysisGUIDRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeGetAnalysisStateByAnalysisGUIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	if err := transportutil.GetContextError(ctx); err != nil {
		wrappedResp := struct {
			Error *transportutil.Error
		}{
			Error: transportutil.MakeError(err),
		}
		w.WriteHeader(wrappedResp.Error.HTTPCode)
		return json.NewEncoder(w).Encode(wrappedResp)
	}

	resp := response.(GetAnalysisStateByAnalysisGUIDResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		GetAnalysisStateByAnalysisGUIDResponse
	}{
		GetAnalysisStateByAnalysisGUIDResponse: resp,
	}

	if resp.err != nil {
		if apierrors.IsErrorLikeResult(resp.err) {
			return transportutil.HandleErrorLikeResult(ctx, w, resp.err)
		}

		terr := transportutil.MakeError(resp.err)
		wrappedResp.Error = terr
		w.WriteHeader(terr.HTTPCode)
	}

	return json.NewEncoder(w).Encode(wrappedResp)
}

func decodeGetAnalysisStateByPRNumberRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetAnalysisStateByPRNumberRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeGetAnalysisStateByPRNumberResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	if err := transportutil.GetContextError(ctx); err != nil {
		wrappedResp := struct {
			Error *transportutil.Error
		}{
			Error: transportutil.MakeError(err),
		}
		w.WriteHeader(wrappedResp.Error.HTTPCode)
		return json.NewEncoder(w).Encode(wrappedResp)
	}

	resp := response.(GetAnalysisStateByPRNumberResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		GetAnalysisStateByPRNumberResponse
	}{
		GetAnalysisStateByPRNumberResponse: resp,
	}

	if resp.err != nil {
		if apierrors.IsErrorLikeResult(resp.err) {
			return transportutil.HandleErrorLikeResult(ctx, w, resp.err)
		}

		terr := transportutil.MakeError(resp.err)
		wrappedResp.Error = terr
		w.WriteHeader(terr.HTTPCode)
	}

	return json.NewEncoder(w).Encode(wrappedResp)
}

func decodeUpdateAnalysisStateByAnalysisGUIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateAnalysisStateByAnalysisGUIDRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeUpdateAnalysisStateByAnalysisGUIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	if err := transportutil.GetContextError(ctx); err != nil {
		wrappedResp := struct {
			Error *transportutil.Error
		}{
			Error: transportutil.MakeError(err),
		}
		w.WriteHeader(wrappedResp.Error.HTTPCode)
		return json.NewEncoder(w).Encode(wrappedResp)
	}

	resp := response.(UpdateAnalysisStateByAnalysisGUIDResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		UpdateAnalysisStateByAnalysisGUIDResponse
	}{
		UpdateAnalysisStateByAnalysisGUIDResponse: resp,
	}

	if resp.err != nil {
		if apierrors.IsErrorLikeResult(resp.err) {
			return transportutil.HandleErrorLikeResult(ctx, w, resp.err)
		}

		terr := transportutil.MakeError(resp.err)
		wrappedResp.Error = terr
		w.WriteHeader(terr.HTTPCode)
	}

	return json.NewEncoder(w).Encode(wrappedResp)
}
