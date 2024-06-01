package response

import (
	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
)

type CreateProduct struct {
	ProductID uuid.UUID `json:"itemId" db:"product_id"`
}

type ListProduct struct {
	ProductID       uuid.UUID                `json:"itemId" db:"product_id"`
	Name            string                   `json:"name" db:"name"`
	ProductCategory constant.ProductCategory `json:"productCategory" db:"product_category"`
	Price           int                      `json:"price" db:"price"`
	ImageUrl        string                   `json:"imageUrl" db:"image_url"`
	CreatedAt       string                   `json:"createdAt" db:"created_at"`
}
