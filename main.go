package main

import (
	"fmt"
	"github.com/anvari1313/yaus/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	c := config.ReadConfig("")
	fmt.Println(c.Server.Addr)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
