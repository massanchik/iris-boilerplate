package main

import (
	"github.com/kataras/iris"
	// Import handlers package, change this to your own package.
	"github.com/iris-contrib/iris-boilerplate/handlers"
)

//Setup routes
func setupMainRoutes(app *iris.Application) {
	app.Handle("GET", "/", handlers.Index)
}
