package domain

import (
	"context"
	"log"
	"strconv"
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

	log.Print(`username -> `, montirAccount.Username)
	log.Print(`password -> `, montirAccount.Password)

	result, err := c.LoginService.MontirLogin(&montirAccount)
	if err != nil {
		return nil, err
	}

	MontirId := strconv.Itoa(int(result.Id))
	return &model.LoginResponeMessage{Message: MontirId + " Berhasil Login"}, nil
}

func (c LoginServer) UserLogin(ctx context.Context, param *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	return &model.LoginResponeMessage{}, nil
}
