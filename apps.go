package main

var mainRoutes = []interface{}{
	Routes["index"],
}

//List of apps(sites)
var Apps = map[string]map[string]interface{}{
	"main-app": map[string]interface{}{
		"redis":  "127.0.0.1:6379", //Some custom config
		"host":   ":8000",          //Resolves to localhost:8000
		"routes": mainRoutes,       //Route
	},
}
