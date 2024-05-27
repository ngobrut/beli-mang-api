package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
)

type Merchant struct {
	MerchantID       uuid.UUID                 `json:"merchantId" db:"merchant_id"`
	Name             string                    `json:"name" db:"name"`
	MerchantCategory constant.MerchantCategory `json:"merchantCategory" db:"merchant_category"`
	ImageUrl         string                    `json:"imageUrl" db:"image_url"`
	Lat              float64                   `json:"lat" db:"lat"`
	Long             float64                   `json:"long" db:"long"`
	CreatedAt        time.Time                 `json:"createdAt" db:"created_at"`
}
