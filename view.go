package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
)

func setupMainView(app *iris.Application) {
	// load templates
	templates := view.Django("./templates", ".html")
	templates.Reload(true) // remove this line if production
	if err := app.AttachView(templates); err != nil {
		panic(err)
	}

	// set favicon or static assets:
	// App.StaticWeb("/static", "./static")
	// App.Favicon("./static/assets/favicon.ico")
}
