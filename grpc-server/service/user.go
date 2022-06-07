package service

import (
	"context"
	api "github.com/eminoz/todo/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *UserCollection) CreateUser(ctx context.Context, customer *api.Customer) (*api.ResCustomer, error) {
	insertOne, err := u.collection.InsertOne(ctx, customer)
	if err != nil {
		panic(err)
	}
	userId := insertOne.InsertedID
	responseCustomer := new(*api.ResCustomer)
	filter := bson.D{{"_id", userId}}
	err = u.collection.FindOne(ctx, filter).Decode(&responseCustomer)
	return *responseCustomer, err
}

func (u *UserCollection) GetUSerById(ctx context.Context, id *api.CustomerId) (*api.ResCustomer, error) {
	objectID, err2 := primitive.ObjectIDFromHex(id.GetId())
	if err2 != nil {
		panic(err2)
	}
	filter := bson.D{{"_id", objectID}}
	responseCustomer := new(*api.ResCustomer)
	err := u.collection.FindOne(ctx, filter).Decode(&responseCustomer)
	return *responseCustomer, err

}
func (u *UserCollection) DeleteUSerById(ctx context.Context, id *api.CustomerId) (*api.ResDeletedUser, error) {
	objectID, err2 := primitive.ObjectIDFromHex(id.Id)
	if err2 != nil {
		panic(err2)
	}

	filter := bson.D{{"_id", objectID}}

	deleteOne, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	if deleteOne.DeletedCount > 0 {
		return &api.ResDeletedUser{
			Email:   "",
			Id:      id.Id,
			Message: "deleted",
		}, err

	}
	return &api.ResDeletedUser{
		Email:   "",
		Id:      id.Id,
		Message: "user bulunamadı",
	}, err
}
