package usecase

import (
	"context"

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
