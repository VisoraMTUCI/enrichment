package config

type Config struct {
	AppName string `env:"APP_NAME"`

	DebugServer   HTTP
	KafkaConusmer Kafka
	CH            Clickhouse
}

func New() {}
