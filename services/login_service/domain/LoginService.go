package domain

import (
	"context"
	"tublessin/common/model"
)

type LoginService struct {
	MontirService model.MontirClient
}

type LoginServiceInterface interface {
	MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
	UserLogin()
}

func NewLoginService(client model.MontirClient) LoginServiceInterface {
	return &LoginService{client}
}

func (s LoginService) MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, error) {
	result, err := s.MontirService.Login(context.Background(), montirAccount)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s LoginService) UserLogin() {}
