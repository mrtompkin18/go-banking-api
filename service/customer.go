package service

type CustomerResponse struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
