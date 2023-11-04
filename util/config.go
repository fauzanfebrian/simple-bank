package util

import (
	"errors"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	EmailSenderHost      string        `mapstructure:"EMAIL_SENDER_HOST"`
	EmailSenderPort      string        `mapstructure:"EMAIL_SENDER_PORT"`
	EmailSenderUsername  string        `mapstructure:"EMAIL_SENDER_USERNAME"`
	GinMode              string
	DBDriver             string
}

func bindEnvKeysToViper(config Config) {
	r := reflect.TypeOf(config)
	for r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	for i := 0; i < r.NumField(); i++ {
		env := r.Field(i).Tag.Get("mapstructure")
		if env != "" {
			viper.BindEnv(env)
		}
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(filePath string) (config Config, err error) {
	viper.SetConfigFile(filePath)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return
		}
		bindEnvKeysToViper(config)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	config.DBDriver = "postgres"

	switch config.Environment {
	case "production":
		config.GinMode = gin.ReleaseMode
	case "test":
		config.GinMode = gin.TestMode
	default:
		config.GinMode = gin.DebugMode
	}

	return
}
