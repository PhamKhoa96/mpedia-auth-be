package config

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
)

type (
	Web struct {
		Debug     bool
		Verbose   bool
		Addr      string
		ImagePath string
	}

	Config struct {
		Web
	}
)

func LoadConfigFromToml(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
