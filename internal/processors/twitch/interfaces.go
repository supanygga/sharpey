package twitch

import (
	"context"

	"github.com/supanygga/sharpey/internal/models/twitch"
	"github.com/supanygga/sharpey/pkg/models/message"
	"github.com/supanygga/sharpey/pkg/models/user"
)

// TwitchClient.
//
//go:generate mockery --name TwitchClient --with-expecter --keeptree --case underscore
type TwitchClient interface {
	// Messages.
	Messages() <-chan *twitch.Message

	// Join.
	Join(ctx context.Context, channels ...string) error

	// SendMessage.
	SendMessage(ctx context.Context, channel, message string) error

	// SendMessagef.
	SendMessagef(ctx context.Context, channel, message string, format ...any) error
}

// UserStorage.
//
//go:generate mockery --name UserStorage --with-expecter --keeptree --case underscore
type UserStorage interface {
	// Create.
	Create(ctx context.Context, req *user.CreateRequest) (*user.CreateResponse, error)

	// Get.
	Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error)

	// Update.
	Update(ctx context.Context, req *user.UpdateRequest) (*user.UpdateResponse, error)

	// Delete.
	Delete(ctx context.Context, req *user.DeleteRequest) (*user.DeleteResponse, error)

	// List.
	List(ctx context.Context, req *user.ListRequest) (*user.ListResponse, error)
}

// MessageStorage.
//
//go:generate mockery --name MessageStorage --with-expecter --keeptree --case underscore
type MessageStorage interface {
	// Create.
	Create(ctx context.Context, req *message.CreateRequest) (*message.CreateResponse, error)

	// Get.
	Get(ctx context.Context, req *message.GetRequest) (*message.GetResponse, error)

	// Update.
	Update(ctx context.Context, req *message.UpdateRequest) (*message.UpdateResponse, error)

	// Delete.
	Delete(ctx context.Context, req *message.DeleteRequest) (*message.DeleteResponse, error)

	// List.
	List(ctx context.Context, req *message.ListRequest) (*message.ListResponse, error)
}
