package bootstrap

import (
	"log"

	"github.com/fatih/color"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	Mode                   string `mapstructure:"MODE"`
}

func NewEnv() *Env {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	env := &Env{}
	err = viper.Unmarshal(env)
	if err != nil {
		log.Fatal("Env can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		color.Green("The App is running in development env")
	}

	return env
}

func (config *Env) Validate() error {
	return validation.ValidateStruct(config,
		validation.Field(&config.AppEnv, validation.Required),
		validation.Field(&config.ServerAddress, validation.Required),
		validation.Field(&config.ContextTimeout, validation.Required),
		validation.Field(&config.DBHost, validation.Required),
		validation.Field(&config.DBPort, is.Port),
		validation.Field(&config.DBUser, validation.Required),
		validation.Field(&config.DBPassword, validation.Required),
		validation.Field(&config.DBName, validation.Required),
		validation.Field(&config.AccessTokenExpiryHour, validation.Required),
		validation.Field(&config.RefreshTokenExpiryHour, validation.Required),
		validation.Field(&config.AccessTokenSecret, validation.Required),
		validation.Field(&config.RefreshTokenSecret, validation.Required),
		validation.Field(&config.Mode, validation.In("debug", "release")),
	)
}
