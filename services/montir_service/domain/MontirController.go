package domain

import (
	"context"
	"database/sql"
	"tublessin/common/model"
)

type MontirServer struct {
	MontirUsecase MontirUsecaseInterface
}

func NewMontirController(db *sql.DB) *MontirServer {
	return &MontirServer{NewMontirUsecase(db)}
}

// Disini adalah pusat Method2 dari Montir-Service
func (c MontirServer) Login(ctx context.Context, param *model.MontirAccount) (*model.MontirAccount, error) {
	result, err := c.MontirUsecase.Login(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c MontirServer) RegisterNewMontir(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	result, err := c.MontirUsecase.RegisterNewMontir(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c MontirServer) UpdateMontirProfileByID(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	return &model.MontirResponeMessage{Response: ""}, nil

}

func (c MontirServer) GetMontirProfileByID(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	return &model.MontirResponeMessage{Response: ""}, nil
}
