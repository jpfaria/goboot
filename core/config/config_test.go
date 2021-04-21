package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DBConfig struct {
	Username string `config:"username"`
}

type MyConfig struct {
	Addr  string
	Debug string
	DB    DBConfig `config:"db"`
	Redis struct {
		Host string `config:"h"`
	} `config:"red"`
}

func TestPFlag(t *testing.T) {

	flagLoad()

	Add("key", "value", "test")

	Load()

	assert.Equal(t, "value", instance.String("key"), "they should be equal")
}

func TestEnv(t *testing.T) {

	flagLoad()

	os.Setenv("K_ENV", "value")

	Load()

	assert.Equal(t, "value", instance.String("k.env"), "they should be equal")
}

func TestConf(t *testing.T) {

	flagLoad()

	os.Args = []string{"--conf", "./testdata/config.json", "--conf", "./testdata/config.yaml"}

	Load()

	assert.True(t, instance.Bool("debug"), "they should be true")
	assert.Equal(t, "127.0.0.13", instance.String("redis.host"), "they should be equal")
}

func TestUnmarshal(t *testing.T) {

	flagLoad()

	var err error

	os.Args = []string{"--conf", "./testdata/config.json", "--conf", "./testdata/config.yaml"}

	Load()

	c := MyConfig{}
	err = Unmarshal(&c)
	assert.Nil(t, err, "they should be nil")
	assert.Equal(t, ":8083", c.Addr, "they should be equal")
	assert.Equal(t, "foosss", c.DB.Username, "they should be equal")
	assert.Equal(t, "127.0.0.14", c.Redis.Host, "they should be equal")
}
