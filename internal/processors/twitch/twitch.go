package twitch

import (
	"context"
	"errors"
	"log/slog"
)

// New.
func New(
	twitchClient TwitchClient,
	userStorage UserStorage,
	messageStorage MessageStorage,
	options *Options,
	logger *slog.Logger,
) (*Processor, error) {
	if twitchClient == nil {
		return nil, errors.New("nil twitch client")
	}

	if userStorage == nil {
		return nil, errors.New("nil user storage")
	}

	if messageStorage == nil {
		return nil, errors.New("nil message storage")
	}

	if options == nil {
		return nil, errors.New("nil options")
	}

	if logger == nil {
		return nil, errors.New("nil logger")
	}

	return &Processor{
		twitchClient:   twitchClient,
		userStorage:    userStorage,
		messageStorage: messageStorage,
		options:        options,
		logger:         logger,
	}, nil
}

// Run.
func (p *Processor) Run(ctx context.Context) error {
	if p.isRunning {
		return errors.New("already running")
	}
	go p.loop(ctx)

	p.isRunning = true

	return nil
}

// loop.
func (p *Processor) loop(ctx context.Context) {
	defer func() {
		p.isRunning = false
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			p.dummy()
		}
	}
}

func (p *Processor) dummy() {}
