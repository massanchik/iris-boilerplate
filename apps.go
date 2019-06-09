package main

import "github.com/kataras/iris"

// Configurators a list of iris.Configurator
// iris.Configurator is just a func(app *iris.Application) so you can
// configure the app and do amazing things, it also accepts a value of *iris.Configuration
// which you can load from *.toml or *.yaml files.
type Configurators []iris.Configurator

// App is our application which contains useful information about our HTTP servers.
type App struct {
	Name string

	// Host contains an iris.Runner which can accept any type of server runner,
	// i.e Addr, TLS, AutomaticTLS, Unix...
	//
	// iris-only magic no1.
	Host iris.Runner
	// Define any type of configurator, remember in go we should accept anything and return specific.
	//
	// iris-only magic no2.
	Config Configurators
}

// Apps is a lis of our apps (sites)
var Apps = []App{
	App{
		Name: "main-app",
		Host: iris.Addr(":8000"), //Resolves to localhost:8000
		Config: Configurators{
			iris.WithOtherValue("redis", "127.0.0.1:6379"), // dynamic custom config goes here
			iris.WithCharset("UTF-8"),
			// iris.WithTray, // so a tray icon in the taskbar to control each instance visually, optionally
			iris.WithoutBanner,           // don't print banner and server information
			iris.WithoutInterruptHandler, // don't catch CTRL+C/CMD+C, we will catch it by ourselves
			// remember? Configurator is just a func(*iris.Application),
			// so our custom functions are valid Configurators:
			setupMainView,
			setupMainRoutes,
		}},
}
