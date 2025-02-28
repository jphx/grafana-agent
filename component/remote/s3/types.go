package s3

import (
	"fmt"
	"time"

	"github.com/grafana/agent/pkg/river"

	"github.com/grafana/agent/pkg/flow/rivertypes"
)

var _ river.Unmarshaler = (*Arguments)(nil)

// Arguments implements the input for the s3 component
type Arguments struct {
	Path string `river:"path,attr"`
	// PollFrequency determines the frequency to check for changes
	// defaults to 10m
	PollFrequency time.Duration `river:"poll_frequency,attr,optional"`
	// IsSecret determines if the content should be displayed to the user
	IsSecret bool `river:"is_secret,attr,optional"`
	// Options allows you to override default settings
	Options ClientOptions `river:"client_options,block,optional"`
}

// ClientOptions implements specific AWS configuration options
type ClientOptions struct {
	AccessKey    string            `river:"key,attr,optional"`
	Secret       rivertypes.Secret `river:"secret,attr,optional"`
	Endpoint     string            `river:"endpoint,attr,optional"`
	DisableSSL   bool              `river:"disable_ssl,attr,optional"`
	UsePathStyle bool              `river:"use_path_style,attr,optional"`
	Region       string            `river:"region,attr,optional"`
}

const minimumPollFrequency = 30 * time.Second

// DefaultArguments sets the poll frequency
var DefaultArguments = Arguments{
	PollFrequency: 10 * time.Minute,
}

// UnmarshalRiver implements the unmarshaller
func (a *Arguments) UnmarshalRiver(f func(v interface{}) error) error {
	*a = DefaultArguments
	type arguments Arguments
	err := f((*arguments)(a))
	if err != nil {
		return err
	}
	if a.PollFrequency <= minimumPollFrequency {
		return fmt.Errorf("poll_frequency must be greater than 30s")
	}
	return nil
}

// Exports implements the file content
type Exports struct {
	Content rivertypes.OptionalSecret `river:"content,attr"`
}
