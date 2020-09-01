package domain

import (
	"database/sql"
	"log"
	"tublessin/common/model"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	Login(username string) (*model.UserAccount, error)
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

// Disini adalah layer Repository dari User-Service, untuk berkomunikasi dengan database
func (r UserRepository) Login(username string) (*model.UserAccount, error) {
	results := r.db.QueryRow("SELECT * FROM user_account WHERE username=?", username)
	var userAccount model.UserAccount

	err := results.Scan(&userAccount.Id, &userAccount.Username, &userAccount.Password, &userAccount.DateCreated)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &userAccount, nil
}
