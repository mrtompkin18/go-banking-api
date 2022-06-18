package service

import "time"

type AccountResponse struct {
	AccountId   int       `json:"account_id"`
	CustomerId  int       `json:"customer_id"`
	Balance     float64   `json:"balance"`
	Status      int       `json:"status"`
	OpeningDate time.Time `json:"opening_date"`
	AccountType string    `json:"account_type"`
}

type AccountRequest struct {
	CustomerId  int     `json:"customer_id"`
	Balance     float64 `json:"balance"`
	AccountType string  `json:"account_type"`
}

type AccountService interface {
	GetAccounts() ([]AccountResponse, error)
	GetAccountByCustomerId(int) (*AccountResponse, error)
	NewAccount(account AccountRequest) (*AccountResponse, error)
}
