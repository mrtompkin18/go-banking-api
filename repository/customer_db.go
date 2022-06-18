package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return customerRepository{db}
}

func (r customerRepository) GetAll() ([]Customer, error) {
	query := "SELECT * FROM customer"

	var customers []Customer
	err := r.db.Select(&customers, query)

	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepository) GetById(id int) (*Customer, error) {
	query := "SELECT * FROM customer where customer_id = ?"

	var customer Customer
	err := r.db.Get(&customer, query, id)

	if err != nil {
		return nil, err
	}
	return &customer, nil
}
