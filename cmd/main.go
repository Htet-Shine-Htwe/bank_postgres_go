package main

import (
	"log"

	"github.com/Htet-Shine-Htwe/bank_go/cmd/api"
	"github.com/Htet-Shine-Htwe/bank_go/db"
)

func main() {

	db, err := db.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000")

	server.Run()
}
