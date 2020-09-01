package domain

import (
	"context"
	"tublessin/common/model"
)

type LoginRepository struct {
	Montir model.MontirClient
}

type LoginRepositoryInterface interface {
	MontirLogin(montirAccount *model.MontirAccount)
	UserLogin()
}

func NewLoginRepository(client model.MontirClient) LoginRepositoryInterface {
	return &LoginRepository{client}
}

func (r LoginRepository) MontirLogin(montirAccount *model.MontirAccount) {
	r.Montir.Login(context.Background(), montirAccount)
}

func (r LoginRepository) UserLogin() {}
