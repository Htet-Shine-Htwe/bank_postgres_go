package main

import "github.com/Htet-Shine-Htwe/bank_go/cmd/api"

func main() {
	server := api.NewAPIServer(":3000")

	server.Run()
}
