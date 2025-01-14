package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/jumayevgadam/evernote-go/internal/app"
	"github.com/jumayevgadam/evernote-go/internal/helpers"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		zap.L().Error("main.godotenv.Load: error", zap.Error(err))
	}

	configPath := flag.String("config", "local", "path to the config file (local or docker)")

	// parse the flags.
	flag.Parse()

	// get actualConfigPath.
	actualConfigPath := helpers.GetConfigPath(*configPath)

	// run application.
	app.Run(actualConfigPath)
}
