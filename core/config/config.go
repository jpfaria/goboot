package config

type Config struct {
	Key         string
	Value       interface{}
	Description string
	Options     *Options
}
