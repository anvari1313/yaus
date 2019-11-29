package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, name string, username string, password string) (string, error)
}

type UserModel struct {
	Name     string `bson:"name"`
	Username string `bson:"user_name"`
	Password string `bson:"pass_word"`
}

type userRepository struct {
	collection *mongo.Collection
}

func CreateUserRepository(database *mongo.Database) *userRepository {
	return &userRepository{
		collection: database.Collection("users"),
	}
}

func (u *userRepository) Create(ctx context.Context, name string, username string, password string) (string, error) {
	user := UserModel{
		Name:     name,
		Username: username,
		Password: password,
	}
	c, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return c.InsertedID.(primitive.ObjectID).Hex(), nil
}
