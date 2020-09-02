package domain

import (
	"context"
	"database/sql"
	"tublessin/common/model"
)

type UserServer struct {
	UserUsecase UserUsecaseInterface
}

func NewUserController(db *sql.DB) *UserServer {
	return &UserServer{NewUserUsecase(db)}
}

// Disini adalah pusat Method2 dari User-Service
func (c UserServer) Login(ctx context.Context, param *model.UserAccount) (*model.UserAccount, error) {
	result, err := c.UserUsecase.Login(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c UserServer) RegisterNewUser(ctx context.Context, param *model.UserAccount) (*model.UserResponeMessage, error) {
	return &model.UserResponeMessage{Response: ""}, nil
}

func (c UserServer) UpdateUserProfileById(ctx context.Context, param *model.UserAccount) (*model.UserResponeMessage, error) {
	return &model.UserResponeMessage{Response: ""}, nil

}

func (c UserServer) GetUserProfileById(ctx context.Context, param *model.UserAccount) (*model.UserResponeMessage, error) {
	return &model.UserResponeMessage{Response: ""}, nil
}
