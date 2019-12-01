package main

import (
	"github.com/anvari1313/yaus/app"
	"github.com/anvari1313/yaus/config"
)

func main() {
	c := config.ReadConfig("")
	app.CreateApp(c)
}
