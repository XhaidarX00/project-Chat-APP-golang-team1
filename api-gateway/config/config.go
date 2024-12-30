package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppDebug        bool
	Email           EmailConfig
	RedisConfig     RedisConfig
	ServerIp        string
	ServerPort      string
	ShutdownTimeout int
}

type RedisConfig struct {
	Url      string
	Password string
	Prefix   string
}

type EmailConfig struct {
	ApiKey    string
	FromName  string
	FromEmail string
}

func LoadConfig() (Config, error) {
	// Set default values
	setDefaultValues()

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")

	// Allow Viper to read environment variables
	viper.AutomaticEnv()

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %s, using default values or environment variables", err)
	}

	// add value to the config
	config := Config{
		AppDebug:        viper.GetBool("APP_DEBUG"),
		ServerIp:        viper.GetString("SERVER_IP"),
		ServerPort:      viper.GetString("SERVER_PORT"),
		ShutdownTimeout: viper.GetInt("SHUTDOWN_TIMEOUT"),

		Email:       loadEmailConfig(),
		RedisConfig: loadRedisConfig(),
	}
	return config, nil
}

func loadRedisConfig() RedisConfig {
	return RedisConfig{
		Url:      viper.GetString("REDIS_URL"),
		Password: viper.GetString("REDIS_PASSWORD"),
		Prefix:   viper.GetString("REDIS_PREFIX"),
	}
}

func loadEmailConfig() EmailConfig {
	return EmailConfig{
		ApiKey:    viper.GetString("MAILERSEND_API_KEY"),
		FromName:  viper.GetString("MAILERSEND_FROM_NAME"),
		FromEmail: viper.GetString("MAILERSEND_FROM_EMAIL"),
	}
}

func setDefaultValues() {
	viper.SetDefault("APP_DEBUG", true)
	viper.SetDefault("SERVER_PORT", "8181")
	viper.SetDefault("SHUTDOWN_TIMEOUT", 5)
}
