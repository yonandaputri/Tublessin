package domain

import (
	"database/sql"
	"tublessin/common/model"
)

type MontirService struct {
	MontirRepository MontirRepositoryInterface
}

type MontirServiceInterface interface {
	Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
}

func NewMontirService(db *sql.DB) MontirServiceInterface {
	return &MontirService{NewMontirRepository(db)}
}

func (s MontirService) Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error) {
	montirDetail, err := s.MontirRepository.Login(montirAccount.Username, "A")
	if err != nil {
		return nil, err
	}

	return montirDetail, nil
}
