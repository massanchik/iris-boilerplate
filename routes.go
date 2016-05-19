package main

//Import handlers package
import (
	"treenode/iris-boilerplate/handlers"
)

//Setup routes
var Routes = map[string]map[string]interface{}{
	"index": map[string]interface{}{
		"handler": handlers.Index, //Handler(controller in other frameworks)
		"path":    "/",            //Relative URL
		"method":  "GET",          //GET, POST, PUT, DELETE or empty string for ANY
	},
}
