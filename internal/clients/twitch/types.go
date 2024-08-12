package twitch

import (
	"log/slog"

	"github.com/ergochat/irc-go/ircevent"
	"github.com/supanygga/sharpey/internal/models/twitch"
)

type Client struct {
	options        *Options
	logger         *slog.Logger
	conn           *ircevent.Connection
	messageChannel chan *twitch.Message
}
