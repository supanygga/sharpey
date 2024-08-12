package converter

import (
	"strings"
	"time"

	"github.com/ergochat/irc-go/ircmsg"
	"github.com/supanygga/sharpey/internal/models/twitch"
)

// IRCToMessage.
func IRCToMessage(e *ircmsg.Message, ts time.Time) *twitch.Message {
	var (
		channel string
		content string
	)

	if len(e.Params) == 2 {
		channel, _ = strings.CutPrefix(e.Params[0], "#")
		content = e.Params[1]
	}

	return &twitch.Message{
		Channel:   channel,
		User:      e.Nick(),
		Content:   content,
		Timestamp: ts,
	}
}
