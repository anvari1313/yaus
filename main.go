package main

import (
	"github.com/anvari1313/yaus/app"
	"github.com/anvari1313/yaus/config"
	"github.com/anvari1313/yaus/db"
	"github.com/anvari1313/yaus/shorturl"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	c := config.ReadConfig("")

	sh, err := shorturl.Create(c.ShortenedURL)
	if err != nil {
		log.Fatalf("can not start application: %s", err.Error())
	}

	redis, err := db.ConnectRedis(c.Redis)
	if err != nil {
		log.Fatalf("can not start application: %s", err.Error())
	}

	mongo, err := db.ConnectMongo(c.MongoDB)

	a := app.CreateApp(mongo, redis, sh)
	e := echo.New()
	e.GET("/:url_id", a.GetURL)
	e.POST("/url", a.PostURL)

	log.Fatal(e.Start(c.Server.Addr))
}
