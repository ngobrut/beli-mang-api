package request

import (
	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
)

type CreateProduct struct {
	Name            string                   `json:"name" validate:"required,min=2,max=30"`
	ProductCategory constant.ProductCategory `json:"productCategory" validate:"required,productCategory"`
	Price           int                      `json:"price" validate:"required,min=1"`
	ImageUrl        string                   `json:"imageUrl" validate:"required,validUrl"`
	MerchantID      uuid.UUID                `json:"-"`
}

type ListProductQuery struct {
	ProductID       *string   `json:"itemId"`
	Limit           *int      `json:"limit"`
	Offset          *int      `json:"offset"`
	Name            *string   `json:"name"`
	ProductCategory *string   `json:"productCategory"`
	CreatedAt       *string   `json:"createdAt"`
	MerchantID      uuid.UUID `json:"-"`
}
