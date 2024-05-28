package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
	"github.com/ngobrut/beli-mang-api/internal/model"
)

// CreateProduct implements IFaceRepository.
func (r *Repository) CreateProduct(ctx context.Context, data *model.Product) error {
	query := `
        INSERT INTO products(merchant_id, name, product_category, price, image_url)
        VALUES (@merchant_id, @name, @product_category, @price, @image_url) RETURNING product_id, created_at`
	args := pgx.NamedArgs{
		"merchant_id":      data.MerchantID,
		"name":             data.Name,
		"product_category": data.ProductCategory,
		"price":            data.Price,
		"image_url":        data.ImageUrl,
	}

	dest := []interface{}{
		&data.ProductID,
		&data.CreatedAt,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" && pgErr.ConstraintName == "product_merchant_fk" {
				return custom_error.SetCustomError(&custom_error.ErrorContext{
					HTTPCode: http.StatusNotFound,
					Message:  constant.HTTPStatusText(http.StatusNotFound),
				})
			}
		}
		return err
	}

	return nil

}
