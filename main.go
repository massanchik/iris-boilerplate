package main

import (
	"os"
	"os/signal"

	"github.com/kataras/iris"
)

//Setup handlers
func main() {

	for _, app := range Apps {
		go func(app App) {
			// create our server
			server := iris.New()

			// create and attach our per-app file based logger, i.e: "main-app.txt"
			logger := newLogFile(app.Name)
			defer logger.Close()
			server.AttachLogger(logger)

			// execute the configurators from our Config, see how easy and structure and de-coupled is it?
			server.Configure(app.Config...)

			// optionally:
			println(app.Name + " started.")

			if err := server.Run(app.Host); err != nil {
				// log any errors or http closed to our file.
				server.Log("'%s': %v", app.Name, err)
				// optionally, print that this server is closed to the terminal too.
				println(app.Name + " closed.")
			}
		}(app)
	}

	// Iris already waits for os.Interrupt or os.Kill
	// to shutdown gracefully the server, but we disabled that feature.
	blocker := make(chan os.Signal, 1)
	signal.Notify(blocker, os.Interrupt, os.Kill)
	select {
	case <-blocker:
		println("exiting...")
	}
}

// cd $GOPATH/src/github.com/massanchik/iris-boilerplate
// go build && ./iris-boilerplate # or iris-boilerplate.exe on Windows
