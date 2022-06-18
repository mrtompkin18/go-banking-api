package repository

import "github.com/jmoiron/sqlx"

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return accountRepository{db}
}

func (r accountRepository) Create(account NewAccount) (*Account, error) {
	query := "INSERT INTO account(customer_id, account_type, balance) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query,
		account.CustomerId,
		account.AccountType,
		account.Balance,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	accountResponse, err := r.GetByAcctId(int(id))
	if err != nil {
		panic(err)
	}

	return accountResponse, nil
}

func (r accountRepository) GetAll() ([]Account, error) {
	query := "SELECT * FROM account"

	var accounts []Account
	err := r.db.Select(&accounts, query)

	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r accountRepository) GetByCustId(customerId int) (*Account, error) {
	query := "SELECT * FROM account where customer_id = ?"

	var account Account
	err := r.db.Get(&account, query, customerId)

	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r accountRepository) GetByAcctId(accountId int) (*Account, error) {
	query := "SELECT * FROM account where account_id = ?"

	var account Account
	err := r.db.Get(&account, query, accountId)

	if err != nil {
		return nil, err
	}
	return &account, nil
}
