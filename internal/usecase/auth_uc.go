package usecase

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
	"github.com/ngobrut/beli-mang-api/internal/model"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
	"github.com/ngobrut/beli-mang-api/util"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error) {
	pwd, err := util.HashPwd(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: pwd,
		Role:     req.Role,
	}

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	claims := &util.CustomClaims{
		UserID: user.UserID.String(),
		Role:   string(user.Role),
	}

	token, err := util.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.AuthResponse{

		Token: token,
	}

	return res, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error) {
	user, err := u.repo.FindOneUserByUsernameAndRole(ctx, req.Username, req.Role)
	if err != nil {
		return nil, err
	}

	err = util.ComparePwd([]byte(user.Password), []byte(req.Password))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "wrong password",
		})

		return nil, err
	}

	claims := &util.CustomClaims{
		UserID: user.UserID.String(),
		Role:   string(user.Role),
	}

	token, err := util.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.AuthResponse{
		Token: token,
	}

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return u.repo.FindOneUserByID(ctx, userID)
}
