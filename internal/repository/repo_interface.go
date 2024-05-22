package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/model"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User) error
	FindOneUserByUsernameAndRole(ctx context.Context, username string, role constant.Role) (*model.User, error)
	FindOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
}
