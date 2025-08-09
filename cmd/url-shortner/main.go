package main

import (
	"log/slog"
	"os"
	"url-shortner/internal/config"
	"url-shortner/internal/lib/logger/setuplogger"
	"url-shortner/internal/lib/logger/sl"
	"url-shortner/internal/storage/sqlite"
)

func main() {

	cfg := config.MustLoad()

	log := setuplogger.Setup(cfg.Env)

	log.Info("starting url-shortner", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	//TODO: init router: chi

	//TODO: run server:

}
