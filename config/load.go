package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Settings Settings `json:"settings"`
	Users    []User   `json:"users"`
}

type Settings struct {
	Address string        `json:"address"`
	Timeout time.Duration `json:"timeout"`
}

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

var config Config

func Load(dirname string) error {
	env := os.Getenv("ENV")
	if env == "" {
		return errors.New("environment variable 'ENV' is not set")
	}

	configFile := fmt.Sprintf("./%s/%s.yml", dirname, env)
	buffer, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	config, err = parseBuffer(buffer)
	if err != nil {
		return err
	}

	return nil
}

func parseBuffer(buffer []byte) (Config, error) {
	data := Config{}

	err := yaml.Unmarshal(buffer, &data)
	if err != nil {
		return Config{}, err
	}

	return data, nil
}
