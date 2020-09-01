package login

import (
	"encoding/json"
	"log"
	"net/http"
	"tublessin/common/model"
)

type LoginControllerApi struct {
	LoginUsecaseApi LoginUsecaseApiInterface
}

func NewLoginControllerApi(loginService model.LoginClient, montirService model.MontirClient, userService model.UserClient) *LoginControllerApi {
	return &LoginControllerApi{LoginUsecaseApi: NewLoginUsecaseApi(loginService, montirService, userService)}
}

// Nangkep request dari depan yang nanti nya akan di teruskan ke Login-Service
// Disini cuman nge parse data json yang masuk
func (c LoginControllerApi) HandleLoginMontir() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirAccount model.MontirLoginForm
		json.NewDecoder(r.Body).Decode(&montirAccount)

		log.Print(`username -> `, montirAccount.Username, "Mencoba Login")

		result, err := c.LoginUsecaseApi.HandleLoginMontir(&montirAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(result)
			return
		}

		log.Print(`Success Login`)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c LoginControllerApi) HandleLoginUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userAccount model.UserLoginForm
		json.NewDecoder(r.Body).Decode(&userAccount)

		log.Print(`username -> `, userAccount.Username, "Mencoba Login")

		result, err := c.LoginUsecaseApi.HandleLoginUser(&userAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(result)
			return
		}

		log.Print(`Success Login`)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c LoginControllerApi) HandleRegisterNewMontir() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirAccount model.MontirAccount
		json.NewDecoder(r.Body).Decode(&montirAccount)

		result, err := c.LoginUsecaseApi.HandleRegisterNewMontir(&montirAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
