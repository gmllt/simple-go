package main

import (
	"fmt"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gmllt/simple-go/logger"
	"github.com/gmllt/simple-go/web"
	"github.com/prometheus/common/version"
)

var (
	PrometheusNamespace = "tjv"
)

func main() {
	panic(boot())
}

func boot() error {
	kingpin.Version(version.Print("simple-go"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	logger.LoadLogConfig(config.Log)

	a := web.NewWeb(config.Web)
	err = a.Serve()
	if err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}
