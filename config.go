package main

import (
	"io"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gmllt/simple-go/logger"
	"github.com/gmllt/simple-go/web"
	"gopkg.in/yaml.v3"
)

var (
	gConfigFile = kingpin.Flag("config", "path to yaml configuration file").Required().File()
)

type ConfigInterface interface {
	Validate() error
}

type Config struct {
	Web *web.Config    `yaml:",inline"`
	Log *logger.Config `yaml:",inline"`
}

func LoadConfig() (*Config, error) {
	config := Config{}
	content, err := io.ReadAll(*gConfigFile)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
