package domain

import (
	"database/sql"
	"log"
	"tublessin/common/model"
)

type MontirRepository struct {
	db *sql.DB
}

type MontirRepositoryInterface interface {
	Login(username, status string) (*model.MontirAccount, error)
	RegisterNewMontir(m *model.MontirAccount) (*model.MontirResponeMessage, error)
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
		return nil, err
	}

	return &montirAccount, nil
}

func (r MontirRepository) RegisterNewMontir(m *model.MontirAccount) (*model.MontirResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt1, _ := tx.Prepare("INSERT INTO montir_account(username, password) VALUE (?,?)")
	stmnt2, _ := tx.Prepare("INSERT INTO montir_profile(montir_account_id, firstname, lastname, gender, city, email, phone_number) VALUE (?,?,?,?,?,?,?)")
	stmnt3, _ := tx.Prepare("INSERT INTO montir_location(montir_account_id) VALUE(?)")
	stmnt4, _ := tx.Prepare("INSERT INTO montir_status(montir_account_id) VALUE(?)")

	result, err := stmnt1.Exec(m.Username, m.Password)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	_, err = stmnt2.Exec(lastInsertID, m.Profile.Firstname, m.Profile.Lastname, m.Profile.Gender, m.Profile.City, m.Profile.Email, m.Profile.PhoneNumber)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	_, err = stmnt3.Exec(lastInsertID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	_, err = stmnt4.Exec(lastInsertID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &model.MontirResponeMessage{Response: "Inserting New Montir Success", Result: m}, nil
}
