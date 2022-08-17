package repository

import (
	"context"
	"customers-example/model"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type CustomerRepositoryInterface interface {
	GetCustomer(ctx context.Context, id int) (model.Customer, error)
	CreateCustomer(ctx context.Context, request model.Customer) (model.Customer, error)
	UpdateCustomer(ctx context.Context, request model.Customer) (model.Customer, error)
	FindCustomers(ctx context.Context, pagination model.PaginationSettings, searchFilter model.SearchFilter) ([]model.Customer, int, error)
	DeleteCustomerById(ctx context.Context, id int64) error
}

type CustomerRepository struct {
	db *pg.DB
}

func NewCustomerRepository(db *pg.DB) CustomerRepositoryInterface {
	return &CustomerRepository{
		db,
	}
}

func (r *CustomerRepository) CreateCustomer(ctx context.Context, request model.Customer) (model.Customer, error) {
	_, err := r.db.Model(&request).Context(ctx).Returning("id").Insert()

	fmt.Println(request)

	if err != nil {
		return request, err
	}
	return request, nil
}

func (r *CustomerRepository) UpdateCustomer(ctx context.Context, request model.Customer) (model.Customer, error) {

	fmt.Println(request)

	u := model.Customer{}
	_, err := r.db.Model(&u).
		Context(ctx).
		Set("first_name = ?", request.FirstName).
		Set("last_name = ?", request.LastName).
		Set("birth_date = ?", request.BirthDate).
		Set("gender = ?", request.Gender).
		Set("email = ?", request.Email).
		Set("address = ?", request.Address).
		Where("id = ?", request.Id).Update()
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *CustomerRepository) GetCustomer(
	ctx context.Context,
	id int) (model.Customer, error) {

	var customer model.Customer

	err := r.db.Model(&customer).Context(ctx).
		Where("id = ?", id).
		Select()

	return customer, err
}

func (r *CustomerRepository) FindCustomers(
	ctx context.Context,
	pagination model.PaginationSettings,
	searchFilter model.SearchFilter) ([]model.Customer, int, error) {

	var customers []model.Customer

	count, err := r.db.Model(&customers).Context(ctx).
		Order(fmt.Sprintf("%s %s", pagination.Sort, pagination.Direction)).
		Where("first_name LIKE ?", fmt.Sprintf("%s%s%s", "%", searchFilter.FirstName, "%")).
		Where("last_name LIKE ?", fmt.Sprintf("%s%s%s", "%", searchFilter.LastName, "%")).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Limit((pagination.Page-1)*pagination.PageSize + pagination.PageSize).
		SelectAndCount()

	if err != nil {
		return customers, 0, err
	}
	return customers, count, err
}

func (r *CustomerRepository) DeleteCustomerById(ctx context.Context, id int64) error {
	c := model.Customer{}
	_, err := r.db.Model(&c).Context(ctx).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
