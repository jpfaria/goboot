package config

type Options struct {
	Hide bool
}

type Option func(options *Options)

func WithHide() Option {
	return func(options *Options) {
		options.Hide = true
	}
}
