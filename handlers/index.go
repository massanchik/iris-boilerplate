package handlers

import (
	"github.com/kataras/iris/context"
)

//Index form "/" (index) route
func Index(ctx context.Context) {
	// template layout
	ctx.ViewLayout("layout.html")
	// template data, or use a custom struct as the second input parameter
	ctx.ViewData("title", "App Title")
	ctx.ViewData("content", "App Content")
	// execute template
	ctx.View("index.html")
}
