package domain

import "tublessin/common/model"

type LoginService struct {
	LoginRepository LoginRepositoryInterface
}

type LoginServiceInterface interface {
	MontirLogin(montirAccount *model.MontirAccount)
	UserLogin()
}

func NewLoginService(client model.MontirClient) LoginServiceInterface {
	return &LoginService{NewLoginRepository(client)}
}

func (s LoginService) MontirLogin(montirAccount *model.MontirAccount) {
	s.LoginRepository.MontirLogin(montirAccount)
}
func (s LoginService) UserLogin() {}
