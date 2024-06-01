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

// CreateProduct implements IFaceUsecase.
func (u *Usecase) CreateProduct(ctx context.Context, req *request.CreateProduct) (*response.CreateProduct, error) {
	product := &model.Product{
		MerchantID:      req.MerchantID,
		Name:            req.Name,
		ProductCategory: req.ProductCategory,
		Price:           req.Price,
		ImageUrl:        req.ImageUrl,
	}

	err := u.repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	res := &response.CreateProduct{
		ProductID: product.ProductID,
	}

	return res, nil
}

// GetListProduct implements IFaceUsecase.
func (u *Usecase) GetListProduct(ctx context.Context, params *request.ListProductQuery) ([]*response.ListProduct, *response.Meta, error) {
	_, err := u.repo.FindOneMerchantByID(ctx, &params.MerchantID)
	if err != nil {
		return nil, nil, err
	}

	res, meta, err := u.repo.FindProducts(ctx, params)
	if err != nil {
		return nil, nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return res, meta, nil

}
