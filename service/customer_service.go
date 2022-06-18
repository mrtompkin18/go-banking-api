package service

import (
	"database/sql"

	"bank-api/errs"
	"bank-api/logger"
	"bank-api/repository"
)

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return customerService{customerRepository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	var customersResponse []CustomerResponse
	for _, customer := range customers {
		customersResponse = append(customersResponse, CustomerResponse{
			Name:    customer.Name,
			City:    customer.City,
			ZipCode: customer.ZipCode,
		})
	}
	return customersResponse, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.repo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	customerResponse := CustomerResponse{
		Name:    customer.Name,
		City:    customer.City,
		ZipCode: customer.ZipCode,
	}
	return &customerResponse, nil

}
