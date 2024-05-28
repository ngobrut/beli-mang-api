package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/beli-mang-api/internal/model"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
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

// GetListMerchant implements IFaceRepository.
func (r *Repository) FindMerchants(ctx context.Context, params *request.ListMerchantQuery) ([]*response.ListMerchant, *response.Meta, error) {
	query := `SELECT 
        (SELECT count(*) FROM merchants) AS cnt,
        merchant_id,
        name,
        merchant_category,
        lat,
        long,
        TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS.US') as created_at
        FROM merchants`

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	meta := &response.Meta{}

	if params.MerchantID != nil {
		if _, err := uuid.Parse(*params.MerchantID); err != nil {
			clause = append(clause, fmt.Sprintf(" merchant_id = $%d", counter))
			args = append(args, uuid.New())
			counter++

		} else {
			clause = append(clause, fmt.Sprintf(" merchant_id = $%d", counter))
			args = append(args, *params.MerchantID)
			counter++
		}

	} else {
		if params.MerchantCategory != nil {
			clause = append(clause, fmt.Sprintf(" phone = $%d", counter))
			args = append(args, *params.MerchantCategory)
			counter++
		}
		if params.Name != nil {
			clause = append(clause, fmt.Sprintf(" name ilike $%d", counter))
			args = append(args, "%"+*params.Name+"%")
			counter++
		}
	}

	if counter > 1 {
		query += " WHERE" + strings.Join(clause, " AND")
	}

	orderClause := " ORDER BY created_at at time zone 'Asia/Jakarta' DESC"
	if params.CreatedAt != nil {
		if *params.CreatedAt == "asc" {
			orderClause = " ORDER BY created_at at time zone 'Asia/Jakarta' ASC"
		}
	}

	query += orderClause

	if params.Limit != nil && *params.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, params.Limit)
		meta.Limit = *params.Limit
		counter++
	} else {
		query += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, 5)
		meta.Limit = 5
		counter++
	}

	if params.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, params.Offset)
		meta.Offset = *params.Offset
		counter++
	} else {
		query += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, 0)
		meta.Offset = 0
		counter++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	res := make([]*response.ListMerchant, 0)
	for rows.Next() {

		m := &response.ListMerchant{}
		err = rows.Scan(
			&meta.Total,
			&m.MerchantID,
			&m.Name,
			&m.MerchantCategory,
			&m.Location.Lat,
			&m.Location.Long,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		res = append(res, m)
	}

	return res, meta, nil

}
