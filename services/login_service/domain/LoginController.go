package domain

import (
	"context"
	"log"
	"tublessin/common/model"
)

type LoginServer struct {
	LoginUsecase LoginUsecaseInterface
}

func NewLoginController(clientMontir model.MontirClient, clientUser model.UserClient) *LoginServer {
	return &LoginServer{NewLoginUsecase(clientMontir, clientUser)}
}

// Disini adalah pusat Method2 dari Login-Service, method2 disini mengacu pada Service login pada file login.proto
func (c LoginServer) MontirLogin(ctx context.Context, param *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	montirAccount := model.MontirAccount{Username: param.Username, Password: param.Password}

	log.Print(`username -> `, montirAccount.Username, "Mencoba Login")

	result, err := c.LoginUsecase.MontirLogin(&montirAccount)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponeMessage{
		Message: "Login Success",
		Token:   "ini nanti di isi token beneran",
		Account: &model.LoginAccountInfo{
			Id:            result.Id,
			Username:      result.Username,
			Password:      result.Password,
			StatusAccount: result.StatusAccount,
		},
	}, nil
}

func (c LoginServer) UserLogin(ctx context.Context, param *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	userAccount := model.UserAccount{Username: param.Username, Password: param.Password}

	log.Print(`username -> `, userAccount.Username, "Mencoba Login")

	result, err := c.LoginUsecase.UserLogin(&userAccount)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponeMessage{
		Message: "Login Success",
		Token:   "ini nanti di isi token beneran",
		Account: &model.LoginAccountInfo{
			Id:            result.Id,
			Username:      result.Username,
			Password:      result.Password,
			StatusAccount: result.DateCreated,
		},
	}, nil
}
