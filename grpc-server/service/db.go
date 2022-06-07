package service

import (
	"github.com/eminoz/todo/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func OrderCollectionSetting() *UserCollection {
	getDatabase := database.GetDatabase()
	return &UserCollection{db: getDatabase, collection: getDatabase.Collection("user")}
}

type TodoCollection struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func TodoCollectionSetting() *TodoCollection {
	getDatabase := database.GetDatabase()
	return &TodoCollection{db: getDatabase, collection: getDatabase.Collection("todo")}
}
