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
	Port              string        `mapstructure:"Port"`
	Mode              string        `mapstructure:"Mode"`
	ReadTimeout       time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout      time.Duration `mapstructure:"WriteTimeout"`
	CtxDefaultTimeout time.Duration `mapstructure:"CtxDefaultTimeout"`
}

// PostgresDB is used for keeping psqlDB params.
type PostgresDB struct {
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	DBName   string `mapstructure:"db_name"`
	SslMode  string `mapstructure:"db_sslmode"`
}

// Redis is used for keeping redis params.
type RedisDB struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}

// Metrics config.
type Metrics struct {
	URL         string `mapstructure:"Url"`
	ServiceName string `mapstructure:"ServiceName"`
}

// Logger config.
type Logger struct {
	Development       bool   `mapstructure:"Development"`
	DisableCaller     bool   `mapstructure:"DisableCaller"`
	DisableStacktrace bool   `mapstructure:"DisableStacktrace"`
	Encoding          string `mapstructure:"Encoding"`
	Level             string `mapstructure:"Level"`
}

// Jaeger config.
type Jaeger struct {
	Host        string `mapstructure:"JaegerHost"`
	ServiceName string `mapstructure:"JaegerServiceName"`
	LogSpans    bool   `mapstructure:"JaegerLogSpans"`
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
