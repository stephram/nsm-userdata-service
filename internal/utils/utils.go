package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(true)

	switch fmt := os.Getenv("LOG_FORMAT"); fmt {
	case "text":
		log.SetFormatter(&log.TextFormatter{
			// DisableColors: false,
			// ForceColors: true,
			FullTimestamp: true,
		})
	case "json":
		setJSONLogFormat()
	default:
		log.Debugf("unknown LOG_FORMAT value: '%s'", fmt)
		setJSONLogFormat()
	}

	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = log.InfoLevel
		log.WithError(err).Debugf("defaulted to %s", level.String())
	}
	log.SetLevel(level)
	log.Printf("log level set to %s", log.GetLevel().String())
}

func setJSONLogFormat() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.Info("set JSON log format")
}

// GetLogger needs to be called once to ensure logrus is configured correctly.
func GetLogger() *log.Logger {
	return log.StandardLogger()
}
