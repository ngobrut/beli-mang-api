package usecase

import (
	"context"

	"github.com/ngobrut/beli-mang-api/internal/model"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
)

// CreateMerchant implements IFaceUsecase.
func (u *Usecase) CreateMerchant(ctx context.Context, req *request.CreateMerchant) (*response.CreateMerchant, error) {
	merchant := &model.Merchant{
		Name:             req.Name,
		MerchantCategory: req.MerchantCategory,
		ImageUrl:         req.ImageUrl,
		Lat:              req.Location.Lat,
		Long:             req.Location.Long,
	}

	err := u.repo.CreateMerchant(ctx, merchant)
	if err != nil {
		return nil, err
	}

	res := &response.CreateMerchant{
		MerchantID: merchant.MerchantID,
	}

	return res, nil

}
