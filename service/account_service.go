package service

import (
	"database/sql"
	"strings"

	"bank-api/errs"
	"bank-api/logger"
	"bank-api/repository"
)

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return accountService{repo}
}

func (s accountService) GetAccounts() ([]AccountResponse, error) {
	accounts, err := s.repo.GetAll()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	var accountResponse []AccountResponse
	for _, account := range accounts {
		accountResponse = append(accountResponse, AccountResponse{
			AccountId:   account.AccountId,
			CustomerId:  account.CustomerId,
			Balance:     account.Balance,
			Status:      account.Status,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
		})
	}

	return accountResponse, nil
}

func (s accountService) GetAccountByCustomerId(id int) (*AccountResponse, error) {
	account, err := s.repo.GetByCustId(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	accountResponse := AccountResponse{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		Balance:     account.Balance,
		Status:      account.Status,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
	}
	return &accountResponse, nil
}

func (s accountService) NewAccount(request AccountRequest) (*AccountResponse, error) {
	if err := validateAccountRequest(request); err != nil {
		return nil, err
	}

	accountRequest := repository.NewAccount{
		CustomerId:  request.CustomerId,
		Balance:     request.Balance,
		AccountType: request.AccountType,
	}

	account, err := s.repo.Create(accountRequest)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	accountResponse := AccountResponse{
		AccountId:   account.AccountId,
		CustomerId:  account.CustomerId,
		Balance:     account.Balance,
		Status:      account.Status,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
	}
	return &accountResponse, nil
}

func validateAccountRequest(request AccountRequest) error {
	if request.Balance < 5000 {
		return errs.NewValidateError("balance must be greater than 5,000")
	} else if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return errs.NewValidateError("account_type should be `saving` or `checking`")
	}
	return nil
}
