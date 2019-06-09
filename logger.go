package main

import (
	"os"
)

func newLogFile(appName string) *os.File {
	filename := appName + ".txt"
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
