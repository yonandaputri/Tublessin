package domain

import (
	"database/sql"
	"errors"
	"log"
	"tublessin/common/model"
)

type MontirRepository struct {
	db *sql.DB
}

type MontirRepositoryInterface interface {
	Login(username, status string) (*model.MontirAccount, error)
}

func NewMontirRepository(db *sql.DB) MontirRepositoryInterface {
	return &MontirRepository{db}
}

// Disini adalah layer Repository dari Montir-Service, untuk berkomunikasi dengan database
func (r MontirRepository) Login(username, status string) (*model.MontirAccount, error) {
	results := r.db.QueryRow("SELECT * FROM montir_account WHERE username=? AND status_account=?", username, status)
	var montirAccount model.MontirAccount

	err := results.Scan(&montirAccount.Id, &montirAccount.Username, &montirAccount.Password, &montirAccount.StatusAccount)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.New("Username atau Password salah")
	}

	return &montirAccount, nil
}
