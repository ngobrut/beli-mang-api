package response

import "github.com/google/uuid"

type CreateMerchant struct {
	MerchantID uuid.UUID `json:"merchantId"`
}
