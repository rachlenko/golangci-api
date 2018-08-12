// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set RepoAnalysisQuerySet

// RepoAnalysisQuerySet is an queryset type for RepoAnalysis
type RepoAnalysisQuerySet struct {
	db *gorm.DB
}

// NewRepoAnalysisQuerySet constructs new RepoAnalysisQuerySet
func NewRepoAnalysisQuerySet(db *gorm.DB) RepoAnalysisQuerySet {
	return RepoAnalysisQuerySet{
		db: db.Model(&RepoAnalysis{}),
	}
}

func (qs RepoAnalysisQuerySet) w(db *gorm.DB) RepoAnalysisQuerySet {
	return NewRepoAnalysisQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) All(ret *[]RepoAnalysis) error {
	return qs.db.Find(ret).Error
}

// AnalysisGUIDEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) AnalysisGUIDEq(analysisGUID string) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("analysis_guid = ?", analysisGUID))
}

// AnalysisGUIDIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) AnalysisGUIDIn(analysisGUID ...string) RepoAnalysisQuerySet {
	if len(analysisGUID) == 0 {
		qs.db.AddError(errors.New("must at least pass one analysisGUID in AnalysisGUIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("analysis_guid IN (?)", analysisGUID))
}

// AnalysisGUIDNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) AnalysisGUIDNe(analysisGUID string) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("analysis_guid != ?", analysisGUID))
}

// AnalysisGUIDNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) AnalysisGUIDNotIn(analysisGUID ...string) RepoAnalysisQuerySet {
	if len(analysisGUID) == 0 {
		qs.db.AddError(errors.New("must at least pass one analysisGUID in AnalysisGUIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("analysis_guid NOT IN (?)", analysisGUID))
}

// CommitSHAEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CommitSHAEq(commitSHA string) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("commit_sha = ?", commitSHA))
}

// CommitSHAIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CommitSHAIn(commitSHA ...string) RepoAnalysisQuerySet {
	if len(commitSHA) == 0 {
		qs.db.AddError(errors.New("must at least pass one commitSHA in CommitSHAIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("commit_sha IN (?)", commitSHA))
}

// CommitSHANe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CommitSHANe(commitSHA string) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("commit_sha != ?", commitSHA))
}

// CommitSHANotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CommitSHANotIn(commitSHA ...string) RepoAnalysisQuerySet {
	if len(commitSHA) == 0 {
		qs.db.AddError(errors.New("must at least pass one commitSHA in CommitSHANotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("commit_sha NOT IN (?)", commitSHA))
}

// Count is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *RepoAnalysis) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CreatedAtEq(createdAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CreatedAtGt(createdAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CreatedAtGte(createdAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CreatedAtLt(createdAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CreatedAtLte(createdAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) CreatedAtNe(createdAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *RepoAnalysis) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) Delete() error {
	return qs.db.Delete(RepoAnalysis{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtEq(deletedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtGt(deletedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtGte(deletedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtIsNotNull() RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtIsNull() RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtLt(deletedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtLte(deletedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) DeletedAtNe(deletedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) GetUpdater() RepoAnalysisUpdater {
	return NewRepoAnalysisUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDEq(ID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDGt(ID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDGte(ID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDIn(ID ...uint) RepoAnalysisQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDLt(ID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDLte(ID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDNe(ID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) IDNotIn(ID ...uint) RepoAnalysisQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) Limit(limit int) RepoAnalysisQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) Offset(offset int) RepoAnalysisQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs RepoAnalysisQuerySet) One(ret *RepoAnalysis) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderAscByCreatedAt() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderAscByDeletedAt() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderAscByID() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByRepoAnalysisStatusID is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderAscByRepoAnalysisStatusID() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("repo_analysis_status_id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderAscByUpdatedAt() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderDescByCreatedAt() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderDescByDeletedAt() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderDescByID() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByRepoAnalysisStatusID is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderDescByRepoAnalysisStatusID() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("repo_analysis_status_id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) OrderDescByUpdatedAt() RepoAnalysisQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadRepoAnalysisStatus is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) PreloadRepoAnalysisStatus() RepoAnalysisQuerySet {
	return qs.w(qs.db.Preload("RepoAnalysisStatus"))
}

// RepoAnalysisStatusIDEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDEq(repoAnalysisStatusID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("repo_analysis_status_id = ?", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDGt(repoAnalysisStatusID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("repo_analysis_status_id > ?", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDGte(repoAnalysisStatusID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("repo_analysis_status_id >= ?", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDIn(repoAnalysisStatusID ...uint) RepoAnalysisQuerySet {
	if len(repoAnalysisStatusID) == 0 {
		qs.db.AddError(errors.New("must at least pass one repoAnalysisStatusID in RepoAnalysisStatusIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("repo_analysis_status_id IN (?)", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDLt(repoAnalysisStatusID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("repo_analysis_status_id < ?", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDLte(repoAnalysisStatusID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("repo_analysis_status_id <= ?", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDNe(repoAnalysisStatusID uint) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("repo_analysis_status_id != ?", repoAnalysisStatusID))
}

// RepoAnalysisStatusIDNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) RepoAnalysisStatusIDNotIn(repoAnalysisStatusID ...uint) RepoAnalysisQuerySet {
	if len(repoAnalysisStatusID) == 0 {
		qs.db.AddError(errors.New("must at least pass one repoAnalysisStatusID in RepoAnalysisStatusIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("repo_analysis_status_id NOT IN (?)", repoAnalysisStatusID))
}

// ResultJSONEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) ResultJSONEq(resultJSON json.RawMessage) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("result_json = ?", resultJSON))
}

// ResultJSONIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) ResultJSONIn(resultJSON ...json.RawMessage) RepoAnalysisQuerySet {
	if len(resultJSON) == 0 {
		qs.db.AddError(errors.New("must at least pass one resultJSON in ResultJSONIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("result_json IN (?)", resultJSON))
}

// ResultJSONNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) ResultJSONNe(resultJSON json.RawMessage) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("result_json != ?", resultJSON))
}

// ResultJSONNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) ResultJSONNotIn(resultJSON ...json.RawMessage) RepoAnalysisQuerySet {
	if len(resultJSON) == 0 {
		qs.db.AddError(errors.New("must at least pass one resultJSON in ResultJSONNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("result_json NOT IN (?)", resultJSON))
}

// SetAnalysisGUID is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetAnalysisGUID(analysisGUID string) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.AnalysisGUID)] = analysisGUID
	return u
}

// SetCommitSHA is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetCommitSHA(commitSHA string) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.CommitSHA)] = commitSHA
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetCreatedAt(createdAt time.Time) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetDeletedAt(deletedAt *time.Time) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetID(ID uint) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.ID)] = ID
	return u
}

// SetRepoAnalysisStatus is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetRepoAnalysisStatus(repoAnalysisStatus RepoAnalysisStatus) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.RepoAnalysisStatus)] = repoAnalysisStatus
	return u
}

// SetRepoAnalysisStatusID is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetRepoAnalysisStatusID(repoAnalysisStatusID uint) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.RepoAnalysisStatusID)] = repoAnalysisStatusID
	return u
}

// SetResultJSON is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetResultJSON(resultJSON json.RawMessage) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.ResultJSON)] = resultJSON
	return u
}

// SetStatus is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetStatus(status string) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.Status)] = status
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) SetUpdatedAt(updatedAt time.Time) RepoAnalysisUpdater {
	u.fields[string(RepoAnalysisDBSchema.UpdatedAt)] = updatedAt
	return u
}

// StatusEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) StatusEq(status string) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("status = ?", status))
}

// StatusIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) StatusIn(status ...string) RepoAnalysisQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status IN (?)", status))
}

// StatusNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) StatusNe(status string) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("status != ?", status))
}

// StatusNotIn is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) StatusNotIn(status ...string) RepoAnalysisQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status NOT IN (?)", status))
}

// Update is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u RepoAnalysisUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) UpdatedAtEq(updatedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) UpdatedAtGt(updatedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) UpdatedAtGte(updatedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) UpdatedAtLt(updatedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) UpdatedAtLte(updatedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs RepoAnalysisQuerySet) UpdatedAtNe(updatedAt time.Time) RepoAnalysisQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set RepoAnalysisQuerySet

// ===== BEGIN of RepoAnalysis modifiers

// RepoAnalysisDBSchemaField describes database schema field. It requires for method 'Update'
type RepoAnalysisDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f RepoAnalysisDBSchemaField) String() string {
	return string(f)
}

// RepoAnalysisDBSchema stores db field names of RepoAnalysis
var RepoAnalysisDBSchema = struct {
	ID                   RepoAnalysisDBSchemaField
	CreatedAt            RepoAnalysisDBSchemaField
	UpdatedAt            RepoAnalysisDBSchemaField
	DeletedAt            RepoAnalysisDBSchemaField
	RepoAnalysisStatusID RepoAnalysisDBSchemaField
	RepoAnalysisStatus   RepoAnalysisDBSchemaField
	AnalysisGUID         RepoAnalysisDBSchemaField
	Status               RepoAnalysisDBSchemaField
	CommitSHA            RepoAnalysisDBSchemaField
	ResultJSON           RepoAnalysisDBSchemaField
}{

	ID:                   RepoAnalysisDBSchemaField("id"),
	CreatedAt:            RepoAnalysisDBSchemaField("created_at"),
	UpdatedAt:            RepoAnalysisDBSchemaField("updated_at"),
	DeletedAt:            RepoAnalysisDBSchemaField("deleted_at"),
	RepoAnalysisStatusID: RepoAnalysisDBSchemaField("repo_analysis_status_id"),
	RepoAnalysisStatus:   RepoAnalysisDBSchemaField("repo_analysis_status"),
	AnalysisGUID:         RepoAnalysisDBSchemaField("analysis_guid"),
	Status:               RepoAnalysisDBSchemaField("status"),
	CommitSHA:            RepoAnalysisDBSchemaField("commit_sha"),
	ResultJSON:           RepoAnalysisDBSchemaField("result_json"),
}

// Update updates RepoAnalysis fields by primary key
// nolint: dupl
func (o *RepoAnalysis) Update(db *gorm.DB, fields ...RepoAnalysisDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                      o.ID,
		"created_at":              o.CreatedAt,
		"updated_at":              o.UpdatedAt,
		"deleted_at":              o.DeletedAt,
		"repo_analysis_status_id": o.RepoAnalysisStatusID,
		"repo_analysis_status":    o.RepoAnalysisStatus,
		"analysis_guid":           o.AnalysisGUID,
		"status":                  o.Status,
		"commit_sha":              o.CommitSHA,
		"result_json":             o.ResultJSON,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update RepoAnalysis %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// RepoAnalysisUpdater is an RepoAnalysis updates manager
type RepoAnalysisUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewRepoAnalysisUpdater creates new RepoAnalysis updater
// nolint: dupl
func NewRepoAnalysisUpdater(db *gorm.DB) RepoAnalysisUpdater {
	return RepoAnalysisUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&RepoAnalysis{}),
	}
}

// ===== END of RepoAnalysis modifiers

// ===== END of all query sets
