package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EnvConfig struct {
	ServerPort      string `mapstructure:"SERVER_PORT"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	MongodbUri      string `mapstructure:"MONGODB_URI"`
	MongodbDatabase string `mapstructure:"MONGODB_DATABASE"`
	// idk what this do
	UseRedis                   bool   `mapstructure:"USE_REDIS"`
	RedisDefaultAddress        string `mapstructure:"REDIS_DEFAULT_ADDRESS"`
	JWTSecretKey               string `mapstructure:"JWT_SECRET"`
	JWTAccessExpirationMinutes int    `mapstructure:"JWT_ACCESS_EXPIRATION_MINUTES"`
	JWTRefreshExpirationDays   int    `mapstructure:"JWT_REFRESH_EXPIRATION_DAYS"`
	Mode                       string `mapstructure:"MODE"`
}

func (config *EnvConfig) Validate() error {
	return validation.ValidateStruct(config,
		validation.Field(&config.ServerPort, is.Port),
		validation.Field(&config.ServerAddress, validation.Required),
		validation.Field(&config.MongodbUri, validation.Required),
		validation.Field(&config.MongodbDatabase, validation.Required),
		validation.Field(&config.UseRedis, validation.Required),
		validation.Field(&config.RedisDefaultAddress),
		validation.Field(&config.JWTSecretKey, validation.Required),
		validation.Field(&config.JWTAccessExpirationMinutes, validation.Required),
		validation.Field(&config.JWTRefreshExpirationDays, validation.Required),
		validation.Field(&config.Mode, validation.In("debug", "release")),
	)
}
