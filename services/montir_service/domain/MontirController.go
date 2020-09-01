package domain

import (
	"context"
	"database/sql"
	"tublessin/common/model"

	"github.com/golang/protobuf/ptypes/empty"
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
	return &model.MontirResponeMessage{Response: ""}, nil
}

func (c MontirServer) UpdateMontirAccount(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	return &model.MontirResponeMessage{Response: ""}, nil

}

func (c MontirServer) UpdateMontirProfile(ctx context.Context, param *model.MontirProfile) (*model.MontirResponeMessage, error) {
	return &model.MontirResponeMessage{Response: ""}, nil

}

func (c MontirServer) GetMontirProfile(ctx context.Context, void *empty.Empty) (*model.MontirProfile, error) {
	return &model.MontirProfile{Firstname: "", Lastname: ""}, nil
}
