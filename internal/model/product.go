package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
)

type Product struct {
	ProductID       uuid.UUID                `json:"itemId" db:"product_id"`
	MerchantID      uuid.UUID                `json:"merchantId" db:"merchant_id"`
	Name            string                   `json:"name" db:"name"`
	ProductCategory constant.ProductCategory `json:"productCategory" db:"product_category"`
	Price           int                      `json:"price" db:"price"`
	ImageUrl        string                   `json:"imageUrl" db:"image_url"`
	CreatedAt       time.Time                `json:"createdAt" db:"created_at"`
}
