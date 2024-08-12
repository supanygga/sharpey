package migrator

// Options.
type Options struct {
	Table    string `yaml:"table" default:"schema_version" validate:"required"`
	Location string `yaml:"location" validate:"required"`
	Version  int32  `yaml:"version" validate:"required"`
}
