package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const EnvVarPrefix = ""

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
}

type LoggerConfig struct {
	DevMode  bool   `envconfig:"DEV_MODE"`
	LogLevel string `envconfig:"LOG_LEVEL"`
	LogIndex string `envconfig:"LOG_INDEX"`

	DockerId string `envconfig:"HOSTNAME"`
	ClsToken string `envconifg:"CLS_TOKEN"`
}

type ServerConfig struct {
	AppAddr             string `envconfig:"APP_ADDR" default:":8080"`
	SystemAddr          string `envconfig:"SYSTEM_ADDR" default:":53000"`
	RedisUrl            string `env:"REDIS_URL" envDefault:"localhost:6379"`
	RedisPassword       string `env:"REDIS_PASSWORD" envDefault:"localhost:6379"`
	RedisDatabase       int    `env:"REDIS_DATABASE" envDefault:"5"`
	RedisMaxIdle        int    `env:"REDIS_MAX_IDLE" envDefault:"5"`
	RedisMaxActive      int    `env:"REDIS_MAX_ACTIVE" envDefault:"0"`
	RedisIdleTimeout    int    `env:"REDIS_IDLE_TIMEOUT" envDefault:"240"`
	RedisConnectTimeout int    `env:"REDIS_CONNECT_TIMEOUT" envDefault:"10000"`
	RedisReadTimeout    int    `env:"REDIS_READ_TIMEOUT" envDefault:"5000"`
	RedisWriteTimeout   int    `env:"REDIS_WRITE_TIMEOUT" envDefault:"5000"`
}

type PostgresConfig struct {
	MasterDSN string `envconfig:"POSTGRES_MASTER_DSN"`

	MaxIdleConns    int           `envconfig:"POSTGRES_MAX_IDLE_CONNS" default:"5"`
	MaxConns        int           `envconfig:"POSTGRES_MAX_CONNS" default:"20"`
	ConnMaxLifetime time.Duration `envconfig:"POSTGRES_CONN_MAX_LIFETIME" default:"10m"`
}

func NewConfigFromEnvVars() (*Config, error) {
	config := &Config{}
	err := config.overrideWithEnvVars()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (config *Config) overrideWithEnvVars() error {
	return envconfig.Process(EnvVarPrefix, config)
}
