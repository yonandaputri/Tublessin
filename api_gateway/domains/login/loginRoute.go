package login

import (
	"log"
	"tublessin/common/config"
	"tublessin/common/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitLoginRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	loginControllerApi := NewLoginControllerApi(connectToServiceLogin())
	subRouter.HandleFunc("/montir", loginControllerApi.HandleLoginMontir()).Methods("POST")
	subRouter.HandleFunc("/user", loginControllerApi.HandleLoginUser()).Methods("POST")
}

// Untuk Connect ke Service-Login
func connectToServiceLogin() model.LoginClient {
	port := config.SERVICE_LOGIN_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Login-Service", port, err)
	}

	return model.NewLoginClient(conn)
}
