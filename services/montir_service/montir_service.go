package main

import (
	"database/sql"
	"log"
	"net"
	"tublessin/common/config"
	"tublessin/common/model"
	"tublessin/services/montir_service/domain"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "admin"
	dbName   = "tublessin_montir"
	dbHost   = "localhost"
	dbPort   = "3306"
)

func main() {
	srv := grpc.NewServer()
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	montirServer := domain.NewMontirController(db)
	model.RegisterMontirServer(srv, montirServer)

	log.Println("Starting Montir-Service server at port", config.SERVICE_MONTIR_PORT)
	l, err := net.Listen("tcp", config.SERVICE_MONTIR_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_MONTIR_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
