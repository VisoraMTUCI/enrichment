package config

type HTTPServer struct {
	Port                    int `env:"HTTP_SERVER_PORT"                       env-default:"8080"`
	ReadHeaderTimeoutSecond int `env:"HTTP_SERVER_READ_HEADER_TIMEOUT_SECOND" env-default:"10"` //nolint:lll
}

type DebugServer struct {
	Port                    int `env:"DEBUG_SERVER_PORT"                      env-default:"8084"`
	ReadHeaderTimeoutSecond int `env:"HTTP_SERVER_READ_HEADER_TIMEOUT_SECOND" env-default:"10"` //nolint:lll
}
