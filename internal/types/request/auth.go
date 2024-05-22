package request

import "github.com/ngobrut/beli-mang-api/constant"

type Register struct {
	Username string        `json:"username" validate:"required,min=5,max=30"`
	Email    string        `json:"email" validate:"required,email"`
	Password string        `json:"password" validate:"required,min=5,max=30"`
	Role     constant.Role `json:"-"`
}

type Login struct {
	Username string        `json:"username" validate:"required,min=5,max=30"`
	Password string        `json:"password" validate:"required,min=5,max=30"`
	Role     constant.Role `json:"-"`
}
