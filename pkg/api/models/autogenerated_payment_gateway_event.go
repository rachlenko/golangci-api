// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set PaymentGatewayEventQuerySet

// PaymentGatewayEventQuerySet is an queryset type for PaymentGatewayEvent
type PaymentGatewayEventQuerySet struct {
	db *gorm.DB
}

// NewPaymentGatewayEventQuerySet constructs new PaymentGatewayEventQuerySet
func NewPaymentGatewayEventQuerySet(db *gorm.DB) PaymentGatewayEventQuerySet {
	return PaymentGatewayEventQuerySet{
		db: db.Model(&PaymentGatewayEvent{}),
	}
}

func (qs PaymentGatewayEventQuerySet) w(db *gorm.DB) PaymentGatewayEventQuerySet {
	return NewPaymentGatewayEventQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) All(ret *[]PaymentGatewayEvent) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *PaymentGatewayEvent) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) CreatedAtEq(createdAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) CreatedAtGt(createdAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) CreatedAtGte(createdAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) CreatedAtLt(createdAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) CreatedAtLte(createdAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) CreatedAtNe(createdAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// DataEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DataEq(data []byte) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("data = ?", data))
}

// DataIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DataIn(data ...[]byte) PaymentGatewayEventQuerySet {
	if len(data) == 0 {
		qs.db.AddError(errors.New("must at least pass one data in DataIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("data IN (?)", data))
}

// DataNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DataNe(data []byte) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("data != ?", data))
}

// DataNotIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DataNotIn(data ...[]byte) PaymentGatewayEventQuerySet {
	if len(data) == 0 {
		qs.db.AddError(errors.New("must at least pass one data in DataNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("data NOT IN (?)", data))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *PaymentGatewayEvent) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) Delete() error {
	return qs.db.Delete(PaymentGatewayEvent{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(PaymentGatewayEvent{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(PaymentGatewayEvent{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtEq(deletedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtGt(deletedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtGte(deletedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtIsNotNull() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtIsNull() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtLt(deletedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtLte(deletedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) DeletedAtNe(deletedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) GetUpdater() PaymentGatewayEventUpdater {
	return NewPaymentGatewayEventUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDEq(ID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDGt(ID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDGte(ID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDIn(ID ...uint) PaymentGatewayEventQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDLt(ID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDLte(ID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDNe(ID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) IDNotIn(ID ...uint) PaymentGatewayEventQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) Limit(limit int) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) Offset(offset int) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs PaymentGatewayEventQuerySet) One(ret *PaymentGatewayEvent) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderAscByCreatedAt() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderAscByDeletedAt() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderAscByID() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderAscByUpdatedAt() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderAscByUserID is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderAscByUserID() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("user_id ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderDescByCreatedAt() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderDescByDeletedAt() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderDescByID() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderDescByUpdatedAt() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrderDescByUserID is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) OrderDescByUserID() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Order("user_id DESC"))
}

// ProviderEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderEq(provider string) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("provider = ?", provider))
}

// ProviderIDEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderIDEq(providerID string) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("provider_id = ?", providerID))
}

// ProviderIDIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderIDIn(providerID ...string) PaymentGatewayEventQuerySet {
	if len(providerID) == 0 {
		qs.db.AddError(errors.New("must at least pass one providerID in ProviderIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("provider_id IN (?)", providerID))
}

// ProviderIDNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderIDNe(providerID string) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("provider_id != ?", providerID))
}

// ProviderIDNotIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderIDNotIn(providerID ...string) PaymentGatewayEventQuerySet {
	if len(providerID) == 0 {
		qs.db.AddError(errors.New("must at least pass one providerID in ProviderIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("provider_id NOT IN (?)", providerID))
}

// ProviderIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderIn(provider ...string) PaymentGatewayEventQuerySet {
	if len(provider) == 0 {
		qs.db.AddError(errors.New("must at least pass one provider in ProviderIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("provider IN (?)", provider))
}

// ProviderNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderNe(provider string) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("provider != ?", provider))
}

// ProviderNotIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) ProviderNotIn(provider ...string) PaymentGatewayEventQuerySet {
	if len(provider) == 0 {
		qs.db.AddError(errors.New("must at least pass one provider in ProviderNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("provider NOT IN (?)", provider))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetCreatedAt(createdAt time.Time) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.CreatedAt)] = createdAt
	return u
}

// SetData is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetData(data []byte) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.Data)] = data
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetDeletedAt(deletedAt *time.Time) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetID(ID uint) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.ID)] = ID
	return u
}

// SetProvider is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetProvider(provider string) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.Provider)] = provider
	return u
}

// SetProviderID is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetProviderID(providerID string) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.ProviderID)] = providerID
	return u
}

// SetType is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetType(typeValue string) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.Type)] = typeValue
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetUpdatedAt(updatedAt time.Time) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUserID is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) SetUserID(userID *uint) PaymentGatewayEventUpdater {
	u.fields[string(PaymentGatewayEventDBSchema.UserID)] = userID
	return u
}

// TypeEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) TypeEq(typeValue string) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("type = ?", typeValue))
}

// TypeIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) TypeIn(typeValue ...string) PaymentGatewayEventQuerySet {
	if len(typeValue) == 0 {
		qs.db.AddError(errors.New("must at least pass one typeValue in TypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("type IN (?)", typeValue))
}

// TypeNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) TypeNe(typeValue string) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("type != ?", typeValue))
}

// TypeNotIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) TypeNotIn(typeValue ...string) PaymentGatewayEventQuerySet {
	if len(typeValue) == 0 {
		qs.db.AddError(errors.New("must at least pass one typeValue in TypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("type NOT IN (?)", typeValue))
}

// Update is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u PaymentGatewayEventUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UpdatedAtEq(updatedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UpdatedAtGt(updatedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UpdatedAtGte(updatedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UpdatedAtLt(updatedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UpdatedAtLte(updatedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UpdatedAtNe(updatedAt time.Time) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UserIDEq is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDEq(userID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id = ?", userID))
}

// UserIDGt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDGt(userID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id > ?", userID))
}

// UserIDGte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDGte(userID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id >= ?", userID))
}

// UserIDIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDIn(userID ...uint) PaymentGatewayEventQuerySet {
	if len(userID) == 0 {
		qs.db.AddError(errors.New("must at least pass one userID in UserIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_id IN (?)", userID))
}

// UserIDIsNotNull is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDIsNotNull() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id IS NOT NULL"))
}

// UserIDIsNull is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDIsNull() PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id IS NULL"))
}

// UserIDLt is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDLt(userID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id < ?", userID))
}

// UserIDLte is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDLte(userID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id <= ?", userID))
}

// UserIDNe is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDNe(userID uint) PaymentGatewayEventQuerySet {
	return qs.w(qs.db.Where("user_id != ?", userID))
}

// UserIDNotIn is an autogenerated method
// nolint: dupl
func (qs PaymentGatewayEventQuerySet) UserIDNotIn(userID ...uint) PaymentGatewayEventQuerySet {
	if len(userID) == 0 {
		qs.db.AddError(errors.New("must at least pass one userID in UserIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_id NOT IN (?)", userID))
}

// ===== END of query set PaymentGatewayEventQuerySet

// ===== BEGIN of PaymentGatewayEvent modifiers

// PaymentGatewayEventDBSchemaField describes database schema field. It requires for method 'Update'
type PaymentGatewayEventDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f PaymentGatewayEventDBSchemaField) String() string {
	return string(f)
}

// PaymentGatewayEventDBSchema stores db field names of PaymentGatewayEvent
var PaymentGatewayEventDBSchema = struct {
	ID         PaymentGatewayEventDBSchemaField
	CreatedAt  PaymentGatewayEventDBSchemaField
	UpdatedAt  PaymentGatewayEventDBSchemaField
	DeletedAt  PaymentGatewayEventDBSchemaField
	Provider   PaymentGatewayEventDBSchemaField
	ProviderID PaymentGatewayEventDBSchemaField
	UserID     PaymentGatewayEventDBSchemaField
	Type       PaymentGatewayEventDBSchemaField
	Data       PaymentGatewayEventDBSchemaField
}{

	ID:         PaymentGatewayEventDBSchemaField("id"),
	CreatedAt:  PaymentGatewayEventDBSchemaField("created_at"),
	UpdatedAt:  PaymentGatewayEventDBSchemaField("updated_at"),
	DeletedAt:  PaymentGatewayEventDBSchemaField("deleted_at"),
	Provider:   PaymentGatewayEventDBSchemaField("provider"),
	ProviderID: PaymentGatewayEventDBSchemaField("provider_id"),
	UserID:     PaymentGatewayEventDBSchemaField("user_id"),
	Type:       PaymentGatewayEventDBSchemaField("type"),
	Data:       PaymentGatewayEventDBSchemaField("data"),
}

// Update updates PaymentGatewayEvent fields by primary key
// nolint: dupl
func (o *PaymentGatewayEvent) Update(db *gorm.DB, fields ...PaymentGatewayEventDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"created_at":  o.CreatedAt,
		"updated_at":  o.UpdatedAt,
		"deleted_at":  o.DeletedAt,
		"provider":    o.Provider,
		"provider_id": o.ProviderID,
		"user_id":     o.UserID,
		"type":        o.Type,
		"data":        o.Data,
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

		return fmt.Errorf("can't update PaymentGatewayEvent %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// PaymentGatewayEventUpdater is an PaymentGatewayEvent updates manager
type PaymentGatewayEventUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewPaymentGatewayEventUpdater creates new PaymentGatewayEvent updater
// nolint: dupl
func NewPaymentGatewayEventUpdater(db *gorm.DB) PaymentGatewayEventUpdater {
	return PaymentGatewayEventUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&PaymentGatewayEvent{}),
	}
}

// ===== END of PaymentGatewayEvent modifiers

// ===== END of all query sets
