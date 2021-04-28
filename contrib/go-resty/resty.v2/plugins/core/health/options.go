package health

import (
	"github.com/b2wdigital/goignite/v2/core/config"
)

type Options struct {
	Name        string
	Host        string
	Endpoint    string
	Enabled     bool
	Description string
	Required    bool
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
