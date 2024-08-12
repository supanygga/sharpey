package twitch

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ergochat/irc-go/ircevent"
	"github.com/supanygga/sharpey/internal/models/twitch"
)

const (
	twitchServer = "irc.chat.twitch.tv:6667"
)

// New.
func New(options *Options, logger *slog.Logger) (*Client, error) {
	if options == nil {
		return nil, errors.New("nil options")
	}

	if logger == nil {
		return nil, errors.New("nil logger")
	}

	return &Client{
		options: options,
		logger:  logger,
		conn: &ircevent.Connection{
			Server:   twitchServer,
			Nick:     options.Name,
			RealName: options.Name,
			Password: options.password(),
			Debug:    options.Debug,
		},
		messageChannel: make(chan *twitch.Message, options.MessageChannelSize),
	}, nil
}

// Run.
func (c *Client) Run(ctx context.Context) error {
	c.conn.AddCallback("PRIVMSG", c.privmsgCallback)

	if err := c.conn.Connect(); err != nil {
		return err
	}

	if err := c.Join(ctx, c.options.Channels...); err != nil {
		return err
	}

	go c.conn.Loop()

	return nil
}

// Join.
func (c *Client) Join(ctx context.Context, channels ...string) error {
	var errs []error

	for _, channel := range channels {
		if err := c.conn.Join(fmt.Sprintf("#%s", channel)); err != nil {
			errs = append(errs, fmt.Errorf("join channel %s:%w", channel, err))
		}
	}

	return errors.Join(errs...)
}

// Messages.
func (c *Client) Messages() <-chan *twitch.Message {
	return c.messageChannel
}

// SendMessagef.
func (c *Client) SendMessagef(ctx context.Context, channel, message string, format ...any) error {
	return c.SendMessage(ctx, channel, fmt.Sprintf(message, format...))
}

// SendMessage.
func (c *Client) SendMessage(ctx context.Context, channel, message string) error {
	if !strings.HasPrefix(message, "#") {
		message = fmt.Sprintf("#%s", message)
	}

	return c.conn.Privmsg(channel, message)
}

// Close.
func (c *Client) Close() error {
	c.logger.Info("closing twitch client")
	c.conn.Quit()
	close(c.messageChannel)
	return nil
}
