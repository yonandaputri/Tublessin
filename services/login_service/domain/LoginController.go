package domain

import (
	"context"
	"tublessin/common/model"
)

type LoginServer struct {
	LoginService LoginServiceInterface
}

func NewLoginController(client model.MontirClient) *LoginServer {
	return &LoginServer{NewLoginService(client)}
}

func (c LoginServer) MontirLogin(ctx context.Context, param *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	montirAccount := model.MontirAccount{Username: param.Username, Password: param.Password}

	c.LoginService.MontirLogin(&montirAccount)
	return &model.LoginResponeMessage{}, nil
}

func (c LoginServer) UserLogin(ctx context.Context, param *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	return &model.LoginResponeMessage{}, nil
}
