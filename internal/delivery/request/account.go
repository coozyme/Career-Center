package request

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/domain/valueobject"
)

type RequestRegister struct {
	Email    string `json:"email"`
	Nama     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewRegisterRequest(req *RequestRegister) *account.AccountDTO {
	return &account.AccountDTO{
		Email:    req.Email,
		Name:     req.Nama,
		Password: req.Password,
	}
}

func NewRegisterByAdminRequest(req *RequestRegister) *account.AccountDTO {
	role := valueobject.NewTypeRolesFromString(req.Role)
	return &account.AccountDTO{
		Email:    req.Email,
		Name:     req.Nama,
		Password: req.Password,
		Role:     role,
	}
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginRequest(req *RequestLogin) *account.AccountDTO {
	return &account.AccountDTO{
		Email:    req.Email,
		Password: req.Password,
	}
}

type RequestPassword struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func NewPassword(req *RequestPassword) *account.UpdatePasswordDTO {
	return &account.UpdatePasswordDTO{
		OldPassword:     req.OldPassword,
		NewPassword:     req.NewPassword,
		ConfirmPassword: req.ConfirmPassword,
	}
}

type RequestForgetPassword struct {
	Email string `json:"email"`
}
