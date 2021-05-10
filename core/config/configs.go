package config

var (
	entries []Config
)

func Add(key string, value interface{}, description string, opts ...Option) {

	o := &Options{}

	for _, opt := range opts {
		opt(o)
	}

	entries = append(entries, Config{
		Key:         key,
		Value:       value,
		Description: description,
		Options:     o,
	})
}

func Entries() []Config {
	return entries
}
