package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
	"github.com/ngobrut/beli-mang-api/internal/model"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(ctx context.Context, data *model.User) error {
	query := `INSERT INTO users(username, email, password, role) VALUES (@username, @email, @password, @role) RETURNING user_id, username, email, role`
	args := pgx.NamedArgs{
		"username": data.Username,
		"email":    data.Email,
		"password": data.Password,
		"role":     data.Role,
	}

	dest := []interface{}{
		&data.UserID,
		&data.Username,
		&data.Email,
		&data.Role,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		if IsDuplicateError(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.HTTPStatusText(http.StatusConflict),
			})
		}

		return err
	}

	return nil
}

// FindOneUserByUsernameAndRole implements IFaceRepository.
func (r *Repository) FindOneUserByUsernameAndRole(ctx context.Context, username string, role constant.Role) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow(ctx, "SELECT * FROM users WHERE username = $1 and role = $2 and deleted_at IS NULL", username, role).
		Scan(
			&res.UserID,
			&res.Username,
			&res.Email,
			&res.Password,
			&res.Role,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, err
	}

	return res, nil
}

// FindOneUserByID implements IFaceRepository.
func (r *Repository) FindOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow(ctx, "SELECT * FROM users WHERE user_id = $1 and deleted_at IS NULL", userID).
		Scan(
			&res.UserID,
			&res.Username,
			&res.Email,
			&res.Password,
			&res.Role,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, err
	}

	return res, nil
}
