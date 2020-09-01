package domain

import (
	"context"
	"database/sql"
	"log"
	"tublessin/common/model"

	"github.com/golang/protobuf/ptypes/empty"
)

type MontirServer struct {
	MontirService MontirServiceInterface
}

func NewMontirController(db *sql.DB) *MontirServer {
	return &MontirServer{NewMontirService(db)}
}

func (c MontirServer) Login(ctx context.Context, param *model.MontirAccount) (*model.MontirAccount, error) {
	log.Print(`username -> `, param.Username)
	log.Print(`password -> `, param.Password)

	result, err := c.MontirService.Login(param)
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
