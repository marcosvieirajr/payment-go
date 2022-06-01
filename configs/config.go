package configs

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HostName        string        `mapstructure:"HOST_NAME"`
	HostPort        string        `mapstructure:"HOST_PORT"`
	GracefulTimeout time.Duration `mapstructure:"GRACEFUL_TIMEOUT"`
	LogLevel        string        `mapstructure:"LOG_LEVEL"`
	DBHost          string        `mapstructure:"DB_HOST"`
	DBPort          int           `mapstructure:"DB_PORT"`
	DBUser          string        `mapstructure:"DB_USER"`
	DBPassword      string        `mapstructure:"DB_PASSWORD"`
	DBName          string        `mapstructure:"DB_NAME"`
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found to loading")
	}

	viper.AutomaticEnv()
	viper.SetDefault("HOST_PORT", "3000")
	viper.SetDefault("GRACEFUL_TIMEOUT", "15s")
	viper.SetDefault("LOG_LEVEL", "info")

	configs := Config{
		HostName:        viper.GetString("HOST_NAME"),
		HostPort:        viper.GetString("HOST_PORT"),
		GracefulTimeout: viper.GetDuration("GRACEFUL_TIMEOUT"),
		LogLevel:        viper.GetString("LOG_LEVEL"),
		DBHost:          viper.GetString("DB_HOST"),
		DBPort:          viper.GetInt("DB_PORT"),
		DBUser:          viper.GetString("DB_USER"),
		DBPassword:      viper.GetString("DB_PASSWORD"),
		DBName:          viper.GetString("DB_NAME"),
	}

	return configs
}
