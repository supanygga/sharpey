package twitch

// Options.
type Options struct {
	// Channels.
	Channels []string `yaml:"channels" validate:"unique,required,dive,required"`
}
