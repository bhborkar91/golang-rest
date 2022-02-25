package users

import (
	"app/src/db"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	client *db.MongoClient
}

func GetUserService(client *db.MongoClient) *UserService {
	return &UserService{client}
}

func (service *UserService) GetUsers() ([]User, error) {
	ctx, cancel, collection := service.client.Collection("users")
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var data []User = []User{}
	if err := cursor.All(service.client.Ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (service *UserService) CreateUser(userData CreateUserDTO) (*User, error) {
	if userData.Name == "InvalidName" {
		return nil, errors.New("Name is invalid")
	}

	newUser := User{
		Id:   primitive.NewObjectID(),
		Name: userData.Name,
	}

	ctx, cancel, collection := service.client.Collection("users")
	defer cancel()

	_, err := collection.InsertOne(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
