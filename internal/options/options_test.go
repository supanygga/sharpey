package options_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supanygga/sharpey/internal/clients/twitch"
	"github.com/supanygga/sharpey/internal/options"
)

func TestNew(t *testing.T) {
	type args struct {
		file string
	}

	type want struct {
		options *options.Options
		wantErr bool
		errMsg  string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				file: "./testdata/ok.yaml",
			},
			want: want{
				options: &options.Options{
					Twitch: &twitch.Options{
						Name:               "name",
						Token:              "token",
						Channels:           []string{"channel", "channel2"},
						Debug:              false,
						MessageChannelSize: 100,
					},
				},
				wantErr: false,
				errMsg:  "",
			},
		},
		{
			name: "invalid no token",
			args: args{
				file: "./testdata/no_token.yaml",
			},
			want: want{
				options: nil,
				wantErr: true,
				errMsg:  "Key: 'Options.Twitch.Token' Error:Field validation for 'Token' failed on the 'required' tag",
			},
		},
		{
			name: "no file",
			args: args{
				file: "./testdata/nonexistent.yaml",
			},
			want: want{
				options: nil,
				wantErr: true,
				errMsg:  "open ./testdata/nonexistent.yaml: no such file or directory",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options, err := options.New(tt.args.file)
			if tt.want.wantErr {
				assert.Nil(t, options)
				assert.EqualError(t, err, tt.want.errMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want.options, options)
			}
		})
	}
}
