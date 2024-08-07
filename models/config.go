package models

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
