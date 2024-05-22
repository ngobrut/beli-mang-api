package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetListNearbyMerchants(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(chi.URLParam(r, "lat"), 32)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	long, err := strconv.ParseFloat(chi.URLParam(r, "long"), 32)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	fmt.Printf("lat: %v\n", lat)
	fmt.Printf("long: %v\n", long)

	// todo:

	h.ResponseOK(w, http.StatusOK, nil)
}
