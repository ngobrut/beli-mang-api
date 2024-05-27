package handler

import (
	"net/http"

	"github.com/ngobrut/beli-mang-api/internal/types/request"
)

func (h *Handler) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var req request.CreateMerchant
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.CreateMerchant(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}
