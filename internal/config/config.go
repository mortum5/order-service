package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Struct to map env values.
type Config struct {
	DBUser      string `mapstructure:"DB_USER"`
	DBPass      string `mapstructure:"DB_PASSWORD"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      int    `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	SwaggerURL  string `mapstructure:"SWAGGER_URL"`
	NatsURL     string `mapstructure:"NATS_URL"`
	ClusterID   string `mapstructure:"CLUSTER_ID"`
	ClientID    string `mapstructure:"CLIENT_ID"`
	TopicName   string `mapstructure:"TOPIC_NAME"`
	DurableName string `mapstructure:"DURABLE_NAME"`
}

// Call to get a new instance of config with .env variables.
func LoadConfig(path string) (config Config, err error) {
	// Set path/location of env file.
	viper.AddConfigPath("../../")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path)
	viper.AddConfigPath("../../config/")

	// Tell viper the name of config file.
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Viper reads all the variables.
	if err = viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("config: read err %v", err)
	}

	// Unmarshal into our struct.
	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("config: unmarshal err %v", err)
	}
	return
}
