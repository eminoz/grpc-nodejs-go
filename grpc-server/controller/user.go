package controller

import (
	"context"
	api "github.com/eminoz/todo/pb"
)

type UserService interface {
	CreateUser(ctx context.Context, customer *api.Customer) (*api.ResCustomer, error)
	GetUSerById(ctx context.Context, id *api.CustomerId) (*api.ResCustomer, error)
	DeleteUSerById(ctx context.Context, id *api.CustomerId) (*api.ResDeletedUser, error)
}
type User struct {
	UserServices UserService
}

func (u User) CreateUser(ctx context.Context, customer *api.Customer) (*api.ResCustomer, error) {
	createUser, err := u.UserServices.CreateUser(ctx, customer)
	if err != nil {
		panic(err)
	}
	return createUser, nil
}

func (u User) GetUser(ctx context.Context, id *api.CustomerId) (*api.ResCustomer, error) {
	getUSerById, err := u.UserServices.GetUSerById(ctx, id)
	if err != nil {
		panic(err)
	}
	return getUSerById, err
}

func (u User) DeleteUSer(ctx context.Context, id *api.CustomerId) (*api.ResDeletedUser, error) {
	deleteUSerById, err := u.UserServices.DeleteUSerById(ctx, id)
	if err != nil {
		panic(err)
	}
	return deleteUSerById, err

}
