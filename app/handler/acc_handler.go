package acc_handler

import (
	"net/http"

	"github.com/Htet-Shine-Htwe/bank_go/app/types"
	"github.com/Htet-Shine-Htwe/bank_go/util"
	"github.com/gorilla/mux"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	pararms := mux.Vars(r)

	account_id := pararms["id"]

	account := types.NewAccount("htet", "shinee")

	return util.WriteJson(w, 200, account)
}
