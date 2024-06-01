package handler

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetListProduct(w http.ResponseWriter, r *http.Request) {
	merchantID, err := uuid.Parse(r.PathValue("merchantId"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  constant.HTTPStatusText(http.StatusNotFound),
		})
		h.ResponseError(w, err)
		return
	}

	qp := r.URL.Query()

	params := &request.ListProductQuery{
		ProductID:       StringPtr(qp.Get("itemId")),
		Name:            StringPtr(qp.Get("name")),
		ProductCategory: StringPtr(qp.Get("productCategory")),
		CreatedAt:       StringPtr(qp.Get("createdAt")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}

	params.MerchantID = merchantID

	res, meta, err := h.uc.GetListProduct(r.Context(), params)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res, meta)
}
