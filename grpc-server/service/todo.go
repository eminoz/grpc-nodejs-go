package service

import (
	"context"
	api "github.com/eminoz/todo/pb"
	"go.mongodb.org/mongo-driver/bson"
)

func (t *TodoCollection) CreateNewTodo(ctx context.Context, todo *api.Todo) (*api.ResTodo, error) {
	insertOne, err := t.collection.InsertOne(ctx, &todo)
	if err != nil {
		panic(err)
	}

	i := new(*api.ResTodo)
	filter := bson.D{{"_id", insertOne.InsertedID}}
	err = t.collection.FindOne(ctx, filter).Decode(&i)

	return *i, err
}
func (t *TodoCollection) AddNewTodo(ctx context.Context, todo *api.Todo) (*api.ResTodo, error) {
	customerId := todo.CustomerId
	filter := bson.D{{"customerid", customerId}}
	for i, _ := range todo.Title {
		update := bson.D{{Key: "$push", Value: bson.D{{"title", todo.Title[i]}}}}
		t.collection.FindOneAndUpdate(ctx, filter, update)
	}
	i := new(*api.ResTodo)
	err := t.collection.FindOne(ctx, filter).Decode(&i)

	return *i, err

}
func (t *TodoCollection) GetTodos(ctx context.Context, id *api.CustomerId) (*api.ResTodo, error) {
	userID := id.GetId()
	filter := bson.D{{"customerid", userID}}
	i := new(*api.ResTodo)
	err := t.collection.FindOne(ctx, filter).Decode(&i)
	if err != nil {
		panic(err)
	}
	return *i, err
}
func (t *TodoCollection) DeleteOneTodo(ctx context.Context, id *api.RemoveTodo) (*api.ResDeleteTodo, error) {
	filter := bson.D{{"customerid", id.CustomerId}}
	update := bson.D{{Key: "$pull", Value: bson.D{{Key: "title", Value: id.Title[0]}}}}
	_, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
	func(ctx context.Context, filter bson.D) {
		var i *api.RemoveTodo
		t.collection.FindOne(ctx, filter).Decode(&i)
		if len(i.Title) == 0 {
			t.collection.DeleteOne(ctx, filter)
		}
	}(ctx, filter)
	return &api.ResDeleteTodo{Title: id.Title[0], Message: "Updated"}, err
}
func (t *TodoCollection) DeleteAllTodos(ctx context.Context, id *api.CustomerId) (*api.ResDeleteTodo, error) {
	filter := bson.D{{"customerid", id.GetId()}}
	_, err := t.collection.DeleteOne(ctx, filter)
	return &api.ResDeleteTodo{Title: "silindi", Message: "Updated"}, err

}
