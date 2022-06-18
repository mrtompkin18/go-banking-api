package repository

import "time"

type NewAccount struct {
	CustomerId  int
	Balance     float64
	AccountType string
}

type Account struct {
	AccountId   int       `db:"account_id"`
	CustomerId  int       `db:"customer_id"`
	OpeningDate time.Time `db:"opening_date"`
	AccountType string    `db:"account_type"`
	Balance     float64   `db:"balance"`
	Status      int       `db:"status"`
}

type AccountRepository interface {
	Create(NewAccount) (*Account, error)
	GetAll() ([]Account, error)
	GetByCustId(int) (*Account, error)
	GetByAcctId(int) (*Account, error)
}
