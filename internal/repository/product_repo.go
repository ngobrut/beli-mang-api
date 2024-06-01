package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
	"github.com/ngobrut/beli-mang-api/internal/model"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
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

// FindProducts implements IFaceRepository.
func (r *Repository) FindProducts(ctx context.Context, params *request.ListProductQuery) ([]*response.ListProduct, *response.Meta, error) {
	query := `SELECT
        (SELECT count(*) FROM products) AS cnt,
        product_id,
        name,
        product_category,
        price,
        image_url,
        TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS.US') as created_at
        FROM products
        WHERE merchant_id = $1
    `

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 2
	args = append(args, params.MerchantID)

	meta := &response.Meta{}

	if params.ProductID != nil {
		if _, err := uuid.Parse(*params.ProductID); err != nil {
			clause = append(clause, fmt.Sprintf(" product_id = $%d", counter))
			args = append(args, uuid.New())
			counter++

		} else {
			clause = append(clause, fmt.Sprintf(" product_id = $%d", counter))
			args = append(args, *params.ProductID)
			counter++
		}

	} else {
		if params.ProductCategory != nil {
			clause = append(clause, fmt.Sprintf(" product_category = $%d", counter))
			args = append(args, *params.ProductCategory)
			counter++
		}
		if params.Name != nil {
			clause = append(clause, fmt.Sprintf(" name ilike $%d", counter))
			args = append(args, "%"+*params.Name+"%")
			counter++
		}
	}

	if counter > 2 {
		query += " AND" + strings.Join(clause, " AND")
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

	res := make([]*response.ListProduct, 0)
	for rows.Next() {

		p := &response.ListProduct{}
		err = rows.Scan(
			&meta.Total,
			&p.ProductID,
			&p.Name,
			&p.ProductCategory,
			&p.Price,
			&p.ImageUrl,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		res = append(res, p)
	}

	return res, meta, nil
}
