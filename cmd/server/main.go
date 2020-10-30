package main

import (
	"context"
	"dotTest/cmd/app/db"
	apisrv "dotTest/internal/app/api"
	storage "dotTest/internal/db"
	service "dotTest/internal/services/api"
	"dotTest/internal/services/jobs"
	desc "dotTest/pkg/api"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load("scripts/.env"); err != nil {
		log.Fatalf("No .env file found %v", err)
	}
}

func main() {
	connAPI := db.Connect()
	connJOB := db.Connect()
	defer db.Close(connJOB)
	defer db.Close(connAPI)
	storageAPI := storage.NewStorage(connAPI)
	storageJOB := storage.NewStorage(connJOB)
	srv := apisrv.NewApi(service.NewService(storageJOB))

	go jobs.Run(storageAPI, context.Background())

	s := grpc.NewServer()

	desc.RegisterDotTestServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("can't listen: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
