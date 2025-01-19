package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/jumayevgadam/evernote-go/internal/app"
	"github.com/jumayevgadam/evernote-go/internal/helpers"
	"go.uber.org/zap"
)

// @title EVERNOTE-GOLANG-GIN
// @version 3.0
// @description This is a simple evernote which written in golang(gin)
// @termsOfService http://swagger.io/terms/

// @contact.name Gadam Jumayev
// @contact.url https://github.com/jumayevgadam
// @contact.email hypergadam@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		zap.Error(err)
	}

	configPath := flag.String("config", "local", "path to the config file (local or docker)")

	// parse the flags.
	flag.Parse()

	// get actualConfigPath.
	actualConfigPath := helpers.GetConfigPath(*configPath)

	// run application.
	app.Run(actualConfigPath)
}
