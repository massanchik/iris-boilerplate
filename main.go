package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

//App configuration
func appConfig(api *iris.Iris) {
	api.Static("/static", "./static/", 1)
	api.Config().Render.Template.Engine = config.PongoEngine
}

//Setup handlers
func main() {
	fmt.Println(iris.Version)
	for _, app := range Apps {
		go func(app map[string]interface{}) {
			api := iris.New()
			appConfig(api)
			for _, route := range app["routes"].([]interface{}) {
				r, ok := route.(map[string]interface{})
				if !ok || r == nil {
					continue
				}
				api.HandleFunc(r["method"].(string), r["path"].(string), r["handler"].(func(*iris.Context)))
			}

			fmt.Println(app["host"])
			api.Listen(app["host"].(string))
		}(app)
	}

	blocker := make(chan bool)
	<-blocker
}
