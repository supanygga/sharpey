package twitch

import "log/slog"

// Processor.
type Processor struct {
	twitchClient   TwitchClient
	userStorage    UserStorage
	messageStorage MessageStorage
	options        *Options
	logger         *slog.Logger
	isRunning      bool
}
