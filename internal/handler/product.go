package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	merchantID, err := uuid.Parse(r.PathValue("merchantId"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  constant.HTTPStatusText(http.StatusNotFound),
		})
		h.ResponseError(w, err)
		return
	}
	var req request.CreateProduct
	err = h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.MerchantID = merchantID

	res, err := h.uc.CreateProduct(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, res, nil)
}
