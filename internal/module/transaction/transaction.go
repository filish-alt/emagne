package transaction

import (
	"context"
	"database/sql"
	"errors"
	"log"

	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/database/models/constant"
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

	ConfirmBySeller(ctx context.Context, id uuid.UUID, sellerEmail string) (db.Transaction, error)
	MarkPaid(ctx context.Context, id uuid.UUID, buyerEmail string) (db.Transaction, error)
	MarkShipped(ctx context.Context, id uuid.UUID, sellerEmail string) (db.Transaction, error)
	MarkDelivered(ctx context.Context, id uuid.UUID, sellerEmail string) (db.Transaction, error)
	StartInspection(ctx context.Context, id uuid.UUID) (db.Transaction, error)
	MarkClosed(ctx context.Context, id uuid.UUID) (db.Transaction, error)
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
		Status:           sql.NullString{String: initialStatus(req.Role), Valid: true},
	}

	tx, err := m.tx.Create(ctx, params)
	if err != nil {
		return db.Transaction{}, err
	}

	// Notify parties
	notifyEmailSMS(req.SellerEmail, sql.NullString{String: req.SellerPhone, Valid: req.SellerPhone != ""}, "New transaction awaiting your confirmation", "A buyer created a transaction and awaits your confirmation.")
	notifyEmailSMS(req.BuyerEmail, sql.NullString{String: req.BuyerPhone, Valid: req.BuyerPhone != ""}, "Transaction created", "Your transaction has been created. Waiting for seller confirmation.")

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
	// Fetch current transaction
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if !isValidTransition(current.Status.String, status.String) {
		return db.Transaction{}, errors.New("invalid transaction status transition")
	}
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

func initialStatus(role constant.Role) string {
	if role == constant.RoleBuyer {
		return string(constant.StatusPendingSellerConfirm)
	}
	return string(constant.StatusAwaitingPayment)
}

func isValidTransition(from, to string) bool {
	switch from {
	case string(constant.StatusPendingSellerConfirm):
		return to == string(constant.StatusAwaitingPayment)
	case string(constant.StatusAwaitingPayment):
		return to == string(constant.StatusPaid)
	case string(constant.StatusPaid):
		return to == string(constant.StatusShipped)
	case string(constant.StatusShipped):
		return to == string(constant.StatusDelivered)
	case string(constant.StatusDelivered):
		return to == string(constant.StatusInspection)
	case string(constant.StatusInspection):
		return to == string(constant.StatusClosed)
	default:
		return false
	}
}

func (m *module) ConfirmBySeller(ctx context.Context, id uuid.UUID, sellerEmail string) (db.Transaction, error) {
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if current.Status.String != string(constant.StatusPendingSellerConfirm) {
		return db.Transaction{}, errors.New("transaction not awaiting seller confirmation")
	}
	if current.SellerEmail != sellerEmail {
		return db.Transaction{}, errors.New("only the seller can confirm this transaction")
	}
	next := sql.NullString{String: string(constant.StatusAwaitingPayment), Valid: true}
	updated, err := m.tx.UpdateStatus(ctx, id, next)
	if err == nil {
		notifyEmailSMS(current.BuyerEmail, current.BuyerPhone, "Seller confirmed", "Seller confirmed the transaction. Proceed to payment.")
	}
	return updated, err
}

func (m *module) MarkPaid(ctx context.Context, id uuid.UUID, buyerEmail string) (db.Transaction, error) {
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if current.Status.String != string(constant.StatusAwaitingPayment) {
		return db.Transaction{}, errors.New("transaction not awaiting payment")
	}
	if current.BuyerEmail != buyerEmail {
		return db.Transaction{}, errors.New("only the buyer can mark as paid")
	}
	next := sql.NullString{String: string(constant.StatusPaid), Valid: true}
	return m.tx.UpdateStatus(ctx, id, next)
}

func (m *module) MarkShipped(ctx context.Context, id uuid.UUID, sellerEmail string) (db.Transaction, error) {
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if current.Status.String != string(constant.StatusPaid) {
		return db.Transaction{}, errors.New("transaction not paid")
	}
	if current.SellerEmail != sellerEmail {
		return db.Transaction{}, errors.New("only the seller can mark shipment")
	}
	next := sql.NullString{String: string(constant.StatusShipped), Valid: true}
	return m.tx.UpdateStatus(ctx, id, next)
}

func (m *module) MarkDelivered(ctx context.Context, id uuid.UUID, sellerEmail string) (db.Transaction, error) {
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if current.Status.String != string(constant.StatusShipped) {
		return db.Transaction{}, errors.New("transaction not shipped")
	}
	if current.SellerEmail != sellerEmail {
		return db.Transaction{}, errors.New("only the seller can mark delivered")
	}
	next := sql.NullString{String: string(constant.StatusDelivered), Valid: true}
	return m.tx.UpdateStatus(ctx, id, next)
}

func (m *module) StartInspection(ctx context.Context, id uuid.UUID) (db.Transaction, error) {
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if current.Status.String != string(constant.StatusDelivered) {
		return db.Transaction{}, errors.New("transaction not delivered")
	}
	next := sql.NullString{String: string(constant.StatusInspection), Valid: true}
	return m.tx.UpdateStatus(ctx, id, next)
}

func (m *module) MarkClosed(ctx context.Context, id uuid.UUID) (db.Transaction, error) {
	current, err := m.tx.GetWithCategory(ctx, id)
	if err != nil {
		return db.Transaction{}, err
	}
	if current.Status.String != string(constant.StatusInspection) {
		return db.Transaction{}, errors.New("transaction not in inspection")
	}
	next := sql.NullString{String: string(constant.StatusClosed), Valid: true}
	return m.tx.UpdateStatus(ctx, id, next)
}

func notifyEmailSMS(email string, phone sql.NullString, subject string, body string) {
	if email != "" {
		log.Printf("EMAIL to %s | %s - %s", email, subject, body)
	}
	if phone.Valid && phone.String != "" {
		log.Printf("SMS to %s | %s", phone.String, body)
	}
}
