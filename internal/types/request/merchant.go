package request

import "github.com/ngobrut/beli-mang-api/constant"

type Location struct {
	Lat  float64 `json:"lat" validate:"required,lat"`
	Long float64 `json:"long" validate:"required,long"`
}

type CreateMerchant struct {
	Name             string                    `json:"name" validate:"required,min=2,max=30"`
	MerchantCategory constant.MerchantCategory `json:"merchantCategory" validate:"required,merchantCategory"`
	ImageUrl         string                    `json:"imageUrl" validate:"required,validUrl"`
	Location         Location                  `json:"location" validate:"required"`
}

type ListMerchantQuery struct {
	MerchantID       *string
	Limit            *int
	Offset           *int
	Name             *string
	MerchantCategory *string
	CreatedAt        *string
}
