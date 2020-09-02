package montir

import (
	"log"
	"tublessin/common/config"
	"tublessin/common/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitMontirRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	montirControllerApi := NewLoginControllerApi(connectToServiceMontir())
	subRouter.HandleFunc("/profile/detail/{id}", montirControllerApi.HandleGetMontirProfileByID()).Methods("GET")
}

// Untuk Connect ke Service-Montir
func connectToServiceMontir() model.MontirClient {
	port := config.SERVICE_MONTIR_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Montir-Service", port, err)
	}

	return model.NewMontirClient(conn)
}
