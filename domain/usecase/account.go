package usecase

import (
	"CareerCenter/domain/entity/account"
	"context"
	"net/http"
)

type UseCaseAccount interface {
	//Usecase Admin
	RegisterAdmin(ctx context.Context, email string, data *account.AccountDTO) error
	//Usecase User
	Register(ctx context.Context, data *account.AccountDTO) error
	Login(ctx context.Context, data *account.AccountDTO) (*http.Cookie, *account.Login, error)
	UpdatePassword(ctx context.Context, email string, password *account.UpdatePasswordDTO) error
	ForgetPassword(ctx context.Context, email string) error
}
