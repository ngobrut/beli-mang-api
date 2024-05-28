package usecase

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/internal/model"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
)

type IFaceUsecase interface {
	// auth
	Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error)
	Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error)

	// image
	UploadImage(ctx context.Context, file *multipart.FileHeader) (*response.ImageResponse, error)

	// merchant
	CreateMerchant(ctx context.Context, req *request.CreateMerchant) (*response.CreateMerchant, error)
	GetListMerchant(ctx context.Context, params *request.ListMerchantQuery) ([]*response.ListMerchant, *response.Meta, error)

	// product
	CreateProduct(ctx context.Context, req *request.CreateProduct) (*response.CreateProduct, error)
}
