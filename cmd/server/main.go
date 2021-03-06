package main

import (
	"log"

	"github.com/HasanShahjahan/go-grpc-services/internal/db"
	"github.com/HasanShahjahan/go-grpc-services/internal/rocket"
)

func Run() error {
	//responsible for initializing and starting
	// our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return nil
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations ")
		return err
	}

	_ := rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
