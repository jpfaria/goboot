package echo

import "github.com/b2wdigital/goignite/v2/core/config"

type Options struct {
	HideBanner bool
	Port       int
	Json       struct {
		Pretty struct {
			Enabled bool
		}
	}
}

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
