package routes

import (
	acc_handler "github.com/Htet-Shine-Htwe/bank_go/app/handler"
	"github.com/Htet-Shine-Htwe/bank_go/util"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/account", util.MakeHTTPHandleFunc(acc_handler.HandleAccount))
	r.HandleFunc("/get-account", util.MakeHTTPHandleFunc(acc_handler.HandleGetAccount))
}
