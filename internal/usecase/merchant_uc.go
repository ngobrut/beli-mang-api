package usecase

import (
	"context"
	"net/http"

	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
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

// GetListMerchant implements IFaceUsecase.
func (u *Usecase) GetListMerchant(ctx context.Context, params *request.ListMerchantQuery) ([]*response.ListMerchant, *response.Meta, error) {
	res, meta, err := u.repo.FindMerchants(ctx, params)
	if err != nil {
		return nil, nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return res, meta, nil
}
