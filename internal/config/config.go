package config

import (
	"log/slog"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/pkg/errors"

	jsoniter "github.com/json-iterator/go"
)

const appName = "harvester"

const (
	envPrefix           = ""
	envTagName          = "env"
	defaultValueTagName = "env-default"
)

type Config struct {
	AppName            string
	LogLevel           string `env:"LOG_LEVEL"           env-default:"DEBUG"`
	NumConsumerWorkers int    `env:"NUM_CONSUMER_WORKER" env-default:"3"`

	PublicServer HTTPServer
	DebugServer  DebugServer

	KafkaConsumer    KafkaConfig
	ClickHouseClient CHConfig
}

func InitConfig() (*Config, error) {
	c := Config{AppName: appName}

	err := env.ParseWithOptions(&c, env.Options{
		TagName:             envTagName,
		DefaultValueTagName: defaultValueTagName,
		Prefix:              envPrefix,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Cannot parse env variables")
	}

	return &c, nil
}

func (x *Config) PrepareLogLevel() slog.Level {
	var level slog.Level = 0

	switch x.LogLevel {
	case "DEBUG":
		level = -4
	case "INFO":
		level = 0
	case "WARN":
		level = 4
	case "ERROR":
		level = 8
	}

	return level
}

func (x *Config) ConfigHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	responseBody, err := jsoniter.Marshal(x)
	if err != nil {
		slog.ErrorContext(request.Context(), errors.Wrap(err, "failed to marshal config").Error())
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = writer.Write(responseBody)
	if err != nil {
		slog.ErrorContext(
			request.Context(),
			errors.Wrap(err, "failed to write config response").Error(),
		)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}
