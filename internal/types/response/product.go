package response

import "github.com/google/uuid"

type CreateProduct struct {
	ProductID uuid.UUID `json:"itemId" db:"product_id"`
}
