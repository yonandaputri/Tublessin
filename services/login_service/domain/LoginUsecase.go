package domain

import (
	"context"
	"tublessin/common/model"
)

type LoginUsecase struct {
	MontirService model.MontirClient
}

type LoginUsecaseInterface interface {
	MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
	UserLogin()
}

func NewLoginUsecase(client model.MontirClient) LoginUsecaseInterface {
	return &LoginUsecase{client}
}

// Karna Login-Service tidak bisa akses langsung ke Database Montir, jadi harus dioper ke Montir-Service
func (s LoginUsecase) MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, error) {
	result, err := s.MontirService.Login(context.Background(), montirAccount)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s LoginUsecase) UserLogin() {}
