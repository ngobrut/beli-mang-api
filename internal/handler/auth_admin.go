package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/types/request"
	"github.com/ngobrut/beli-mang-api/util"
)

func (h *Handler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var req request.Register
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.Role = constant.AdminRole

	res, err := h.uc.Register(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, res, nil)
}

func (h *Handler) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var req request.Login
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.Role = constant.AdminRole

	res, err := h.uc.Login(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res, nil)
}

func (h *Handler) GetProfileAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := uuid.Parse(util.GetUserIDFromCtx(ctx))
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.GetProfile(ctx, userID)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res, nil)
}
