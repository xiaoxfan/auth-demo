package config

import (
	"github.com/go-yaml/yaml"
	"os"
)

type t struct {
	Server struct {
		Port string
	}
	Jwt struct {
		SigningString string
	}
	Database struct {
		Driver   string
		Username string
		Password string
		Url      string
	}
}

func (t t) DBUrl() string {
	return t.Database.Username + ":" + t.Database.Password + "@" + t.Database.Url
}

var Config t

func init() {
	Config = loadConfig()
}
func loadConfig() t {
	file, err := os.Open("../application.yml")
	if err != nil {
		panic(err)
	}
	var config t
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}
