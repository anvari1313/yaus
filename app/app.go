package app

import (
	"context"
	"fmt"
	"github.com/anvari1313/yaus/shorturl"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type App struct {
	db     *mongo.Database
	redis  *redis.Client
	hashID *shorturl.HashID
}

func CreateApp(db *mongo.Database, redis *redis.Client, hashID *shorturl.HashID) *App {
	return &App{
		db:     db,
		redis:  redis,
		hashID: hashID,
	}
}

type URLModel struct {
	ID       primitive.ObjectID `bson:"_id"`
	ShortUrl string             `bson:"short_url"`
	URL      string             `bson:"url"`
}

func (a *App) GetURL(ctx echo.Context) error {
	urlID := ctx.Param("url_id")
	log.Debugf("Get url with id: `%s`", urlID)

	dbCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res := a.db.Collection("urls").FindOne(dbCtx, bson.M{"short_url": urlID})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Debugf("No document found error with url id: `%s`", urlID)
			return echo.ErrNotFound
		}

		log.Debugf("Error in fetching url from database with id: `%s`", urlID)
		return err
	}

	var url URLModel
	err := res.Decode(&url)
	if err != nil {

		return err
	}

	fmt.Println(url)

	return ctx.Redirect(http.StatusFound, url.URL)
}

type AddUrlBody struct {
	URL string `json:"url"`
}

type AddUrlSuccess struct {
	URL string `json:"url"`
}

func (a *App) PostURL(ctx echo.Context) error {
	body := new(AddUrlBody)
	err := ctx.Bind(body)
	if err != nil {
		return err
	}

	cmd := a.redis.Incr("counter")
	v, err := cmd.Result()
	if err != nil {
		return err
	}

	val, err := a.hashID.Generate(v)
	if err != nil {
		return err
	}

	dbCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	dbRes, err := a.db.
		Collection("urls").
		InsertOne(dbCtx,
			bson.M{
				"short_url": val,
				"url":       body.URL,
			})
	if err != nil {
		return err
	}

	log.Debugf("Insert a new document with id: `%s`", dbRes.InsertedID)

	return ctx.JSON(http.StatusOK, AddUrlSuccess{URL: val})
}
