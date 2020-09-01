package login

import (
	"encoding/json"
	"log"
	"net/http"
	"tublessin/common/model"
)

type LoginControllerApi struct {
	LoginServiceApi LoginServiceApiInterface
}

func NewLoginControllerApi(loginService model.LoginClient) *LoginControllerApi {
	return &LoginControllerApi{LoginServiceApi: NewLoginServiceApi(loginService)}
}

// Nangkep request dari depan yang nanti nya akan di teruskan ke Login-Service
// Disini cuman nge parse data json yang masuk
func (c LoginControllerApi) HandleLoginMontir() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirAccount model.MontirLoginForm
		json.NewDecoder(r.Body).Decode(&montirAccount)

		log.Print(`username -> `, montirAccount.Username)
		log.Print(`password -> `, montirAccount.Password)

		result, err := c.LoginServiceApi.HandleLoginMontir(&montirAccount)
		if err != nil {
			log.Print(`gagal login`)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		log.Print(`Success Login`)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c LoginControllerApi) HandleLoginUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
