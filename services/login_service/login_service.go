package main

import (
	"log"
	"net"
	"tublessin/common/config"
	"tublessin/common/model"
	"tublessin/services/login_service/domain"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()

	loginServer := domain.NewLoginController(connectToServiceMontir(), connectToServiceUser())
	model.RegisterLoginServer(srv, loginServer)

	log.Println("Starting Login-Service server at port", config.SERVICE_LOGIN_PORT)
	l, err := net.Listen("tcp", config.SERVICE_LOGIN_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_LOGIN_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}

func connectToServiceMontir() model.MontirClient {
	port := config.SERVICE_MONTIR_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Montir-Service", port, err)
	}

	return model.NewMontirClient(conn)
}

func connectToServiceUser() model.UserClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to User-Service", port, err)
	}

	return model.NewUserClient(conn)
}
