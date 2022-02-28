package users

import (
	"app/src/common"
	"app/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	client *db.MongoClient
}

func GetUserService(client *db.MongoClient) *UserService {
	return &UserService{client}
}

func (service *UserService) GetUsers() ([]User, common.IAppError) {
	ctx, cancel, collection := service.client.Collection("users")
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, common.ServerError("Internal server error").WithCause(err)
	}

	var data []User = []User{}
	if err := cursor.All(service.client.Ctx, &data); err != nil {
		return nil, common.ServerError("Internal server error").WithCause(err)
	}

	return data, nil
}

func (service *UserService) CreateUser(userData CreateUserDTO) (*User, common.IAppError) {
	if userData.Name == "InvalidNamePanic" {
		panic("Name is invalid")
		// return nil, errors.New("Name is invalid")
	}

	if userData.Name == "InvalidName" {
		return nil, common.BadRequest("The name is invalid")
	}

	newUser := User{
		Id:   primitive.NewObjectID(),
		Name: userData.Name,
	}

	ctx, cancel, collection := service.client.Collection("users")
	defer cancel()

	_, err := collection.InsertOne(ctx, newUser)

	if err != nil {
		return nil, common.ServerError("Failed to create new user").WithCause(err)
	}

	return &newUser, nil
}
