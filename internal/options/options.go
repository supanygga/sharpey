package options

import (
	"io"
	"os"

	"github.com/creasty/defaults"
	"github.com/supanygga/sharpey/internal/validator"
	"gopkg.in/yaml.v3"
)

// New.
func New(file string) (*Options, error) {
	var (
		options Options
		err     error
	)

	yamlFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()

	content, err := io.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(content, &options); err != nil {
		return nil, err
	}

	if err = defaults.Set(&options); err != nil {
		return nil, err
	}

	if err = validator.V.Struct(&options); err != nil {
		return nil, err
	}

	return &options, nil
}
