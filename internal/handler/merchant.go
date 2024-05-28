package handler

import (
	"net/http"
	"strconv"

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

	h.ResponseOK(w, http.StatusCreated, res, nil)
}

func (h *Handler) GetListMerchant(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListMerchantQuery{
		MerchantID:       StringPtr(qp.Get("merchantId")),
		Name:             StringPtr(qp.Get("name")),
		MerchantCategory: StringPtr(qp.Get("merchantCategory")),
		CreatedAt:        StringPtr(qp.Get("createdAt")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}

	res, meta, err := h.uc.GetListMerchant(r.Context(), params)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res, meta)

}
