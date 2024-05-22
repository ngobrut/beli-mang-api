package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ngobrut/beli-mang-api/config"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/middleware"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
	"github.com/ngobrut/beli-mang-api/internal/usecase"
)

func InitHTTPHandler(cnf config.Config, uc usecase.IFaceUsecase) http.Handler {
	h := Handler{
		uc: uc,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestLogger)
	r.Use(middleware.Recover)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: "Error",
			Error: &response.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Success: true,
			Message: "Success",
			Data: map[string]interface{}{
				"app-name": "beli-mang-api",
			},
		})
	})

	r.Route("/", func(r chi.Router) {
		r.Route("/admin", func(admin chi.Router) {
			admin.Post("/register", h.RegisterAdmin)
			admin.Post("/login", h.LoginAdmin)

			admin.Group(func(private chi.Router) {
				private.Use(middleware.Authorize(cnf.JWTSecret, constant.AdminRole))
				private.Route("/profile", func(profile chi.Router) {
					profile.Get("/", h.GetProfileAdmin)
				})

				// private.Route("/merchants", func(merchant chi.Router) {
				// 	merchant.Post("/", h.CreateMerchant)
				// 	merchant.Get("/", h.GetListMerchant)

				// 	merchant.Route("/{merchantID}/items", func(item chi.Router) {
				// 		item.Post("/", h.CreateMerchantItem)
				// 		item.Get("/", h.GetListMerchantItem)
				// 	})
				// })
			})
		})

		r.Route("/users", func(user chi.Router) {
			user.Post("/register", h.RegisterUser)
			user.Post("/login", h.LoginUser)

			user.Group(func(private chi.Router) {
				private.Use(middleware.Authorize(cnf.JWTSecret, constant.UserRole))
				private.Route("/profile", func(profile chi.Router) {
					profile.Get("/", h.GetProfileUser)
				})

				// private.Post("/estimate", h.CalculateOrderEstimation)

				// private.Route("/orders", func(order chi.Router) {
				// 	order.Post("/", h.CreateOrder)
				// 	order.Get("/", h.GetListOrder)
				// })
			})
		})

		r.Group(func(user chi.Router) {
			// user.Use(middleware.Authorize(cnf.JWTSecret, constant.UserRole))
			user.Get("/merchants/nearby/{lat},{long}", h.GetListNearbyMerchants)
		})

		r.Group(func(admin chi.Router) {
			admin.Use(middleware.Authorize(cnf.JWTSecret, constant.AdminRole))
			admin.Post("/image", h.UploadImage)
		})
	})

	return r
}
