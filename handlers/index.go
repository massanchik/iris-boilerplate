package handlers

import (
	"github.com/kataras/iris"
)

//Handler form "/"(index) route
func Index(c *iris.Context) {
	c.Render("index.html", map[string]interface{}{
		"title": "App Title",
		"content": "App Content",
	})
}
