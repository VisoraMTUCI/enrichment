package config

import "time"

type KafkaConfig struct {
	Brokers     []string      `env:"KAFKA_BROKERS"      envSeparator:","`
	CursorTopic string        `env:"KAFKA_CURSOR_TOPIC"`
	DialTimeout time.Duration `env:"KAFKA_DIAL_TIMEOUT"                  env-default:"10s"`
}
