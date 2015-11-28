package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/error"
)

func main() {
	app := f.CreateApp()
	app.Set("env", "prod")
	app.Use(error.Create("Something went wrong!"))
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		panic("panic")
	})
	app.Listen(3000)
}
