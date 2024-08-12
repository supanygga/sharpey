package twitch

import "fmt"

// Options.
type Options struct {
	Name               string   `yaml:"name" validate:"required"`
	Token              string   `yaml:"token" validate:"required"`
	Channels           []string `yaml:"channels" validate:"unique,required,dive,required"`
	Debug              bool     `yaml:"debug"`
	MessageChannelSize int      `yaml:"message_channel_size" default:"100"`
}

// password.
func (o *Options) password() string {
	return fmt.Sprintf("oauth:%s", o.Token)
}
