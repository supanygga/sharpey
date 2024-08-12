package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/supanygga/application"
	"github.com/supanygga/sharpey/internal/clients/twitch"
	"github.com/supanygga/sharpey/internal/options"
	proc "github.com/supanygga/sharpey/internal/processors/twitch"
	"github.com/supanygga/sharpey/internal/storage/postgres"
)

func main() {
	var app *application.Application
	logger := slog.Default()
	ctx, cancel := context.WithCancel(context.Background())

	onStart := func() {
		var err error

		options, err := options.New("./config.yaml")
		if err != nil {
			logAndExit(logger, "parse options", err)
		}

		twitchClient, err := twitch.New(options.Twitch, logger)
		if err != nil {
			logAndExit(logger, "create twitch client", err)
		}
		if err = twitchClient.Run(ctx); err != nil {
			logAndExit(logger, "run twitch client", err)
		}

		store, err := postgres.New(ctx, options.Postgres, logger)
		if err != nil {
			logAndExit(logger, "create store", err)
		}

		processor, err := proc.New(twitchClient, store.User(), store.Message(), options.Processor, logger)
		if err != nil {
			logAndExit(logger, "create twitch processor", err)
		}
		if err = processor.Run(ctx); err != nil {
			logAndExit(logger, "run twitch processor", err)
		}

		app.AddClosers(twitchClient, store)
	}

	onShutdown := func() {
		cancel()
		app.ExecuteClosers()
	}

	app = application.New(onStart, onShutdown, "sharpey", logger)
	app.Run()
}

func logAndExit(logger *slog.Logger, msg string, err error) {
	logger.Error(msg, "error", err)
	os.Exit(1)
}
