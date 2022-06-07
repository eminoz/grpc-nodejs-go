package controller

import (
	"context"
	api "github.com/eminoz/todo/pb"
)

type TodoService interface {
	CreateNewTodo(ctx context.Context, todo *api.Todo) (*api.ResTodo, error)
	AddNewTodo(ctx context.Context, todo *api.Todo) (*api.ResTodo, error)
	GetTodos(ctx context.Context, id *api.CustomerId) (*api.ResTodo, error)
	DeleteOneTodo(ctx context.Context, id *api.RemoveTodo) (*api.ResDeleteTodo, error)
	DeleteAllTodos(ctx context.Context, id *api.CustomerId) (*api.ResDeleteTodo, error)
}
type Todo struct {
	TodoService TodoService
}

func (t Todo) CreateTodo(ctx context.Context, todo *api.Todo) (*api.ResTodo, error) {
	newTodo, err := t.TodoService.CreateNewTodo(ctx, todo)
	if err != nil {
		panic(err)
	}

	return newTodo, nil
}

func (t Todo) AddNewToto(ctx context.Context, todo *api.Todo) (*api.ResTodo, error) {
	addNewTodo, err := t.TodoService.AddNewTodo(ctx, todo)
	if err != nil {
		panic(err)
	}
	return addNewTodo, err
}

func (t Todo) GetTodo(ctx context.Context, id *api.CustomerId) (*api.ResTodo, error) {
	todos, err := t.TodoService.GetTodos(ctx, id)
	if err != nil {
		panic(err)
	}
	return todos, err
}

func (t Todo) DeleteTodo(ctx context.Context, id *api.RemoveTodo) (*api.ResDeleteTodo, error) {
	deleteOneTodo, err := t.TodoService.DeleteOneTodo(ctx, id)
	if err != nil {
		panic(err)
	}
	return deleteOneTodo, err
}

func (t Todo) DeleteAllTodo(ctx context.Context, id *api.CustomerId) (*api.ResDeleteTodo, error) {
	todos, err := t.TodoService.DeleteAllTodos(ctx, id)
	if err != nil {
		panic(err)
	}
	return todos, err
}
