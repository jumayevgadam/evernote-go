package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config struct keeps needed configuration for this project.
type Config struct {
	Server   Server
	Postgres PostgresDB
	Redis    RedisDB
	Logger   Logger
	Metrics  Metrics
	Jaeger   Jaeger
}

// Server is used for keeping server params.
type Server struct {
	Port              string
	Mode              string
	JwtSecretKey      string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
}

// PostgresDB is used for keeping psqlDB params.
type PostgresDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

// Redis is used for keeping redis params.
type RedisDB struct {
	Address  string
	Password string
}

// Metrics config.
type Metrics struct {
	URL         string
	ServiceName string
}

// Logger config.
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Jaeger config.
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

// LoadConfig func is used for loading configuration.
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}

		return nil, errors.New("error reading config file")
	}

	return v, nil
}

// ParseConfig file is used for parsing configuration.
func ParseConfig(v *viper.Viper) (*Config, error) {
	if v == nil {
		return nil, fmt.Errorf("viper instance nil")
	}

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &c, nil
}
