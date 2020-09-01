package login

import (
	"context"
	"tublessin/common/model"
)

type LoginUsecaseApi struct {
	LoginService model.LoginClient
}

type LoginUsecaseApiInterface interface {
	HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error)
}

func NewLoginUsecaseApi(loginService model.LoginClient) LoginUsecaseApiInterface {
	return LoginUsecaseApi{LoginService: loginService}
}

// Dioper ke Login-Service untuk ditangani ada di folder services/login_services/domain/LoginController.go
func (s LoginUsecaseApi) HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	result, err := s.LoginService.MontirLogin(context.Background(), montirAccount)
	if err != nil {
		return &model.LoginResponeMessage{Message: montirAccount.Username, Token: "Username atau Password Salah"}, err
	}

	return result, nil
}
