package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/beli-mang-api/internal/model"
)

// CreateMerchant implements IFaceRepository.
func (r *Repository) CreateMerchant(ctx context.Context, data *model.Merchant) error {
	query := `INSERT INTO merchants(name, merchant_category, image_url, lat, long)
        VALUES (@name, @merchant_category, @image_url, @lat, @long) RETURNING merchant_id, created_at`
	args := pgx.NamedArgs{
		"name":              data.Name,
		"merchant_category": data.MerchantCategory,
		"image_url":         data.ImageUrl,
		"lat":               data.Lat,
		"long":              data.Long,
	}

	dest := []interface{}{
		&data.MerchantID,
		&data.CreatedAt,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		return err
	}

	return nil

}
