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
