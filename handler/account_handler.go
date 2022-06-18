package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bank-api/errs"
	"bank-api/logger"
	"bank-api/service"
	"github.com/gorilla/mux"
)

type accountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) accountHandler {
	return accountHandler{service}
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accountResponse, err := h.service.GetAccounts()
	if err != nil {
		logger.Log.Error("cannot get account list: " + err.Error())
		errs.HandleError(w, err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accountResponse)
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["customerId"])

	if r.Header.Get("content-type") != "application/json" {
		logger.Log.Error("content-type must be support application/json")
		errs.HandleError(w, errs.NewValidateError("request body is invalid format"))
		return
	}

	var accountRequest service.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&accountRequest)
	if err != nil {
		logger.Log.Error("cannot decode request body to entity, reason: " + err.Error())
		errs.HandleError(w, errs.NewValidateError("request body is invalid format"))
		return
	}

	accountRequest.CustomerId = id
	accountResponse, err := h.service.NewAccount(accountRequest)
	if err != nil {
		errs.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accountResponse)
}

func (h accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["customerId"])
	accountResponse, err := h.service.GetAccountByCustomerId(id)

	if err != nil {
		errs.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accountResponse)
}
