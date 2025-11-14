package dto

import (
	"github.com/filagot/emagne/internal/database/models/constant"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateTransaction struct {
	Title            string          `json:"title"`
	Role             constant.Role   `json:"role"`
	Currency         string          `json:"currency"`
	InspectionPeriod string          `json:"inspection_period"`
	ItemCategoryId   uuid.UUID       `json:"item_catagory_id"`
	ItemName         string          `json:"item_name"`
	ItemDescription  string          `json:"item_description"`
	Price            decimal.Decimal `json:"price"`
	ShippingMethod   string          `json:"shipping_method"`
	SellerEmail      string          `json:"seller_email"`
	SellerPhone      string          `json:"seller_phone"`
	BuyerEmail       string          `json:"buyer_email"`
	BuyerPhone       string          `json:"buyer_phone"`
}
