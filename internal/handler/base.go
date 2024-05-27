package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ngobrut/beli-mang-api/constant"
	"github.com/ngobrut/beli-mang-api/internal/custom_error"
	"github.com/ngobrut/beli-mang-api/internal/types/response"
	"github.com/ngobrut/beli-mang-api/internal/usecase"
)

type ValidatorError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ValidatorError) Error() string {
	return e.Message
}

type Handler struct {
	uc usecase.IFaceUsecase
}

func (h Handler) ValidateStruct(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println("[error-parse-body]", err.Error())
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "please check your body request",
		})

		return err
	}

	validate := validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterValidation("merchantCategory", validateMerchantCategory)
	validate.RegisterValidation("validUrl", validateURL)
	validate.RegisterValidation("lat", validateLat)
	validate.RegisterValidation("long", validateLong)

	err = validate.Struct(data)
	if err == nil {
		return nil
	}

	var message string
	var details = make([]string, 0)
	for _, field := range err.(validator.ValidationErrors) {
		message = field.Translate(trans)
		switch field.Tag() {
		case "merchantCategory":
			message = fmt.Sprintf("%s must be one of [%s]", field.Field(), strings.Join(constant.MerchantCategories, ", "))
		case "validUrl":
			message = "image_url should be url"
		case "lat":
			message = "lat should be in latitude format"
		case "long":
			message = "long should be in longitude format"
		}
		details = append(details, message)
	}

	err = ValidatorError{
		Code:    http.StatusBadRequest,
		Message: "request doesn’t pass validation",
		Details: details,
	}

	return err
}

func validateMerchantCategory(fl validator.FieldLevel) bool {
	return constant.ValidMerchantCategory[fl.Field().String()]
}

func validateURL(fl validator.FieldLevel) bool {
	parsedURL, err := url.Parse(fl.Field().String())
	if err != nil {
		return false
	}

	// Check if the scheme is present and it's http or https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Check if the host is present and it has a valid format
	if parsedURL.Host == "" {
		return false
	}

	// Check if the host has a valid domain format
	parts := strings.Split(parsedURL.Host, ".")
	if len(parts) < 2 {
		return false
	}

	// Check if the path, if present, is in a valid format
	if parsedURL.Path != "" && !strings.HasPrefix(parsedURL.Path, "/") {
		return false
	}

	// All checks passed, URL is valid
	return true
}

func validateLat(fl validator.FieldLevel) bool {
	lat := fl.Field().Float()
	return lat >= -90 && lat <= 90
}

func validateLong(fl validator.FieldLevel) bool {
	long := fl.Field().Float()
	return long >= -180 && long <= 180
}

func (h Handler) ResponseOK(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response.JsonResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

func (h Handler) ResponseError(w http.ResponseWriter, err error) {
	v, isValidationErr := err.(ValidatorError)
	if isValidationErr {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: "ValidationError",
			Error: &response.ErrorResponse{
				Code:    v.Code,
				Message: v.Message,
				Details: v.Details,
			},
		})

		return
	}

	e, isCustomErr := err.(*custom_error.CustomError)
	if !isCustomErr {
		if err != nil && !errors.Is(err, context.Canceled) {
			fmt.Println(err.Error(), "[unhandled-error]")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Error: &response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: constant.HTTPStatusText(http.StatusInternalServerError),
			},
		})

		return
	}

	httpCode := http.StatusInternalServerError
	msg := constant.HTTPStatusText(httpCode)

	if e.ErrorContext != nil && e.ErrorContext.HTTPCode > 0 {
		httpCode = e.ErrorContext.HTTPCode
		msg = constant.HTTPStatusText(httpCode)

		if e.ErrorContext.Message != "" {
			msg = e.ErrorContext.Message
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response.JsonResponse{
		Message: http.StatusText(httpCode),
		Error: &response.ErrorResponse{
			Code:    httpCode,
			Message: msg,
		},
	})
}
