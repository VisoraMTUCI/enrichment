package config

type CHConfig struct {
	Host     string `env:"CLICKHOUSE_HOST" env-default:"clickhouse"`
	Port     int    `env:"CLICKHOUSE_PORT" env-default:"9000"`
	User     string `env:"CLICKHOUSE_USER" env-default:"visora-ch-user"`
	Database string `env:"CLICKHOUSE_DATABASE" env-default:"visora_metrcis"`
}
