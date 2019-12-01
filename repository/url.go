package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepository interface {
	Create(ctx context.Context, url string) (string, error)
	FindByID(ctx context.Context, id string) (*URLModel, error)
}

type URLModel struct {
	URL string `bson:"name"`
}

type urlRepository struct {
	collection *mongo.Collection
}

func CreateURLRepository(database *mongo.Database) URLRepository {
	return &urlRepository{collection: database.Collection("urls")}
}

func (u *urlRepository) Create(ctx context.Context, url string) (string, error) {
	m := URLModel{URL: url}
	c, err := u.collection.InsertOne(ctx, m)
	if err != nil {
		return "", err
	}
	return c.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *urlRepository) FindByID(ctx context.Context, id string) (*URLModel, error) {
	res := new(URLModel)
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	rawRes := u.collection.FindOne(ctx, bson.M{"_id": _id})
	err = rawRes.Err()
	if err != nil {
		return nil, err
	}

	err = rawRes.Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
