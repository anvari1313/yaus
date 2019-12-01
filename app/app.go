package app

import (
	"github.com/anvari1313/yaus/util"
	"log"
	"net/http"

	"github.com/anvari1313/yaus/config"
	"github.com/anvari1313/yaus/db"
	"github.com/anvari1313/yaus/repository"
	"github.com/labstack/echo/v4"
)

type App struct {
	UserRepo repository.UserRepository
	URLRepo  repository.URLRepository
	JWTGen   util.JWTGenerator
}

func CreateApp(c *config.Config) {
	mongoDB, err := db.MongoConnect(c.MongoDB)
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		UserRepo: repository.CreateUserRepository(mongoDB),
		URLRepo:  repository.CreateURLRepository(mongoDB),
		JWTGen:   util.CreateJWTGenerator(c.JWT),
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/:id", app.GetRoute)
	e.POST("/register", app.RegisterRoute)

	e.Logger.Fatal(e.Start(c.Server.Addr))
}
