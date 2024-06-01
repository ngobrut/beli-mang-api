package response

import (
	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
)

type CreateMerchant struct {
	MerchantID uuid.UUID `json:"merchantId"`
}

type Location struct {
	Lat  float64 `json:"lat" db:"lat"`
	Long float64 `json:"long" db:"long"`
}

type ListMerchant struct {
	MerchantID       uuid.UUID                 `json:"merchantId" db:"merchant_id"`
	Name             string                    `json:"name" db:"name"`
	MerchantCategory constant.MerchantCategory `json:"merchantCategory" db:"merchant_category"`
	ImageUrl         string                    `json:"imageUrl" db:"image_url"`
	Location         Location                  `json:"location" db:"-"`
	CreatedAt        string                    `json:"createdAt" db:"created_at"`
}
