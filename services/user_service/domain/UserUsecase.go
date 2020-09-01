package domain

import (
	"database/sql"
	"tublessin/common/model"
)

type UserUsecase struct {
	UserRepository UserRepositoryInterface
}

type UserUsecaseInterface interface {
	Login(UserAccount *model.UserAccount) (*model.UserAccount, error)
}

func NewUserUsecase(db *sql.DB) UserUsecaseInterface {
	return &UserUsecase{NewUserRepository(db)}
}

// Ini Adalah Layer Service dari User-Service, untuk menangani bussiness logic
func (s UserUsecase) Login(UserAccount *model.UserAccount) (*model.UserAccount, error) {
	userDetail, err := s.UserRepository.Login(UserAccount.Username)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}
