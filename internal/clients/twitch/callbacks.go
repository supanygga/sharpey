package twitch

import (
	"time"

	"github.com/ergochat/irc-go/ircmsg"
	"github.com/supanygga/sharpey/internal/models/converter"
)

// privmsgCallback.
func (c *Client) privmsgCallback(e ircmsg.Message) {
	go func() {
		c.messageChannel <- converter.IRCToMessage(&e, time.Now())
	}()
}
