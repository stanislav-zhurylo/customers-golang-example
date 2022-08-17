package service

import (
	"context"
	"customers-example/model"
	"customers-example/repository"
)

type CustomerServiceInterface interface {
	GetCustomer(ctx context.Context, id int) (model.Customer, error)
	CreateCustomer(ctx context.Context, request model.Customer) (model.Customer, error)
	UpdateCustomer(ctx context.Context, request model.Customer) (model.Customer, error)
	FindCustomers(ctx context.Context, pagination model.PaginationSettings, searchFilter model.SearchFilter) ([]model.Customer, int, error)
	DeleteCustomerById(ctx context.Context, id int64) error
}

type CustomerService struct {
	customerRepository repository.CustomerRepositoryInterface
}

func NewCustomerService(customerRepository repository.CustomerRepositoryInterface) CustomerServiceInterface {
	return &CustomerService{
		customerRepository,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, request model.Customer) (model.Customer, error) {
	return s.customerRepository.CreateCustomer(ctx, request)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, request model.Customer) (model.Customer, error) {
	return s.customerRepository.UpdateCustomer(ctx, request)
}

func (s *CustomerService) GetCustomer(
	ctx context.Context,
	id int) (model.Customer, error) {

	return s.customerRepository.GetCustomer(ctx, id)
}

func (s *CustomerService) FindCustomers(
	ctx context.Context,
	pagination model.PaginationSettings,
	searchFilter model.SearchFilter) ([]model.Customer, int, error) {

	return s.customerRepository.FindCustomers(ctx, pagination, searchFilter)
}

func (s *CustomerService) DeleteCustomerById(
	ctx context.Context,
	id int64) error {

	return s.customerRepository.DeleteCustomerById(ctx, id)
}
