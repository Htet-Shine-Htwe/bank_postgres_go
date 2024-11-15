package api

import (
	"log"
	"net/http"

	"github.com/Htet-Shine-Htwe/bank_go/routes"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	log.Println("Api Server was serving at localhost", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}
