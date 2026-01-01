package transaction

import (
	"context"
	"database/sql"

	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/filagot/emagne/internal/storage"
	"github.com/google/uuid"
)

type Module interface {
	Create(ctx context.Context, req *dto.CreateTransaction) (db.Transaction, error)
	GetWithCategory(ctx context.Context, id uuid.UUID) (db.GetTransactionWithCategoryRow, error)
	ListByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset int32) ([]db.ListTransactionsByCategoryRow, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status sql.NullString) (db.Transaction, error)

	AddAttribute(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID, value string) (db.TransactionItemAttribute, error)
	ListAttributes(ctx context.Context, txID uuid.UUID) ([]db.ListTransactionItemAttributesRow, error)
	GetAttribute(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID) (db.TransactionItemAttribute, error)
	DeleteAllAttributes(ctx context.Context, txID uuid.UUID) error
}

type module struct {
	tx     storage.TransactionStorage
	txAttr storage.TransactionItemAttributeStorage
}

func New(tx storage.TransactionStorage, txAttr storage.TransactionItemAttributeStorage) Module {
	return &module{tx: tx, txAttr: txAttr}
}

func (m *module) Create(ctx context.Context, req *dto.CreateTransaction) (db.Transaction, error) {
	params := db.CreateTransactionParams{
		Title:            sql.NullString{String: req.Title, Valid: req.Title != ""},
		Role:             db.Role(req.Role),
		Currency:         req.Currency,
		InspectionPeriod: sql.NullString{String: req.InspectionPeriod, Valid: req.InspectionPeriod != ""},
		ItemCategoryID:   req.ItemCategoryId,
		ItemName:         req.ItemName,
		ItemDescription:  sql.NullString{String: req.ItemDescription, Valid: req.ItemDescription != ""},
		Price:            req.Price,
		ShippingMethod:   sql.NullString{String: req.ShippingMethod, Valid: req.ShippingMethod != ""},
		SellerEmail:      req.SellerEmail,
		SellerPhone:      sql.NullString{String: req.SellerPhone, Valid: req.SellerPhone != ""},
		BuyerEmail:       req.BuyerEmail,
		BuyerPhone:       sql.NullString{String: req.BuyerPhone, Valid: req.BuyerPhone != ""},
		Status:           sql.NullString{String: "Pending", Valid: true},
	}

	tx, err := m.tx.Create(ctx, params)
	if err != nil {
		return db.Transaction{}, err
	}

	for _, attr := range req.Attributes {
		_, err := m.txAttr.Insert(ctx, tx.ID, attr.AttributeID, attr.Value)
		if err != nil {
			// Rollback: delete the created transaction (and attributes via cascade or manual delete if needed)
			// Assuming CASCADE delete is set up on FK, otherwise calling DeleteAllAttributes first is safer.
			// Just to be safe, we delete attributes first if any were inserted (though if insert failed, some might be there)
			_ = m.txAttr.DeleteAll(ctx, tx.ID)
			_ = m.tx.Delete(ctx, tx.ID)
			return db.Transaction{}, err
		}
	}

	return tx, nil
}

func (m *module) GetWithCategory(ctx context.Context, id uuid.UUID) (db.GetTransactionWithCategoryRow, error) {
	return m.tx.GetWithCategory(ctx, id)
}

func (m *module) ListByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset int32) ([]db.ListTransactionsByCategoryRow, error) {
	return m.tx.ListByCategory(ctx, categoryID, limit, offset)
}

func (m *module) UpdateStatus(ctx context.Context, id uuid.UUID, status sql.NullString) (db.Transaction, error) {
	return m.tx.UpdateStatus(ctx, id, status)
}

func (m *module) AddAttribute(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID, value string) (db.TransactionItemAttribute, error) {
	return m.txAttr.Insert(ctx, txID, attributeID, value)
}

func (m *module) ListAttributes(ctx context.Context, txID uuid.UUID) ([]db.ListTransactionItemAttributesRow, error) {
	return m.txAttr.List(ctx, txID)
}

func (m *module) GetAttribute(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID) (db.TransactionItemAttribute, error) {
	return m.txAttr.Get(ctx, txID, attributeID)
}

func (m *module) DeleteAllAttributes(ctx context.Context, txID uuid.UUID) error {
	return m.txAttr.DeleteAll(ctx, txID)
}
