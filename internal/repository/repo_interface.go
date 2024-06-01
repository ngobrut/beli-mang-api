package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/model"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User) error
	FindOneUserByUsernameAndRole(ctx context.Context, username string, role constant.Role) (*model.User, error)
	FindOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)

	// merchant
	CreateMerchant(ctx context.Context, data *model.Merchant) error
	FindOneMerchantByID(ctx context.Context, ID *uuid.UUID) (*model.Merchant, error)
	FindMerchants(ctx context.Context, params *request.ListMerchantQuery) ([]*response.ListMerchant, *response.Meta, error)

	// product
	CreateProduct(ctx context.Context, data *model.Product) error
	FindProducts(ctx context.Context, params *request.ListProductQuery) ([]*response.ListProduct, *response.Meta, error)
}
