package logger

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func LoadLogConfig(config *Config) {
	if config.Level == "" {
		return
	}
	switch strings.ToUpper(config.Level) {
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
		return
	case "WARN":
		log.SetLevel(log.WarnLevel)
		return
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
		return
	case "PANIC":
		log.SetLevel(log.PanicLevel)
		return
	case "FATAL":
		log.SetLevel(log.FatalLevel)
		return
	}
	if config.JSON != nil {
		log.Debugf("Log JSON: %v", *config.JSON)
		if *config.JSON {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{
				DisableColors: config.NoColor,
			})
		}
	}
}
