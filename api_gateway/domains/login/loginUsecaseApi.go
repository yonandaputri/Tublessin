package login

import (
	"context"
	"log"
	"tublessin/common/model"
)

type LoginUsecaseApi struct {
	LoginService model.LoginClient
}

type LoginUsecaseApiInterface interface {
	HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error)
	HandleLoginUser(userAccount *model.UserLoginForm) (*model.LoginResponeMessage, error)
}

func NewLoginUsecaseApi(loginService model.LoginClient) LoginUsecaseApiInterface {
	return LoginUsecaseApi{LoginService: loginService}
}

// Dioper ke Login-Service untuk ditangani ada di folder services/login_services/domain/LoginController.go
func (s LoginUsecaseApi) HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	result, err := s.LoginService.MontirLogin(context.Background(), montirAccount)
	if err != nil {
		log.Print(err.Error())
		return &model.LoginResponeMessage{
			Message: "Username atau Password Salah",
			Token:   "0",
			Account: nil,
		}, err
	}

	return result, nil
}

func (s LoginUsecaseApi) HandleLoginUser(userAccount *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	result, err := s.LoginService.UserLogin(context.Background(), userAccount)
	if err != nil {
		log.Print(err.Error())
		return &model.LoginResponeMessage{
			Message: "Username atau Password Salah",
			Token:   "0",
			Account: nil,
		}, err
	}

	return result, nil
}
