package domain

import (
	"context"
	"log"
	"strconv"
	"tublessin/common/model"
)

type LoginServer struct {
	LoginUsecase LoginUsecaseInterface
}

func NewLoginController(client model.MontirClient) *LoginServer {
	return &LoginServer{NewLoginUsecase(client)}
}

// Disini adalah pusat Method2 dari Login-Service, method2 disini mengacu pada Service login pada file login.proto
func (c LoginServer) MontirLogin(ctx context.Context, param *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	montirAccount := model.MontirAccount{Username: param.Username, Password: param.Password}

	log.Print(`username -> `, montirAccount.Username)
	log.Print(`password -> `, montirAccount.Password)

	result, err := c.LoginUsecase.MontirLogin(&montirAccount)
	if err != nil {
		return nil, err
	}

	MontirId := strconv.Itoa(int(result.Id))
	return &model.LoginResponeMessage{Message: MontirId, Token: "asdasdqweqweq123123123"}, nil
}

func (c LoginServer) UserLogin(ctx context.Context, param *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	return &model.LoginResponeMessage{}, nil
}
