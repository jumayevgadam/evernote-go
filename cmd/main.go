package main

import (
	"flag"

	"github.com/jumayevgadam/evernote-go/internal/app"
	"github.com/jumayevgadam/evernote-go/internal/helpers"
)

func main() {
	configPath := flag.String("config", "local", "path to the config file (local or docker)")

	// parse the flags.
	flag.Parse()

	// get actualConfigPath.
	actualConfigPath := helpers.GetConfigPath(*configPath)

	// run application.
	app.Run(actualConfigPath)
}
