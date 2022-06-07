package config

import (
	"github.com/joho/godotenv"
	"os"
)

var Config *Configuration

type Configuration struct {
	MongoDb string `mapstructure:"MONGODB_URI"`
	Port    string `mapstructure:"PORT"`
}

func SetupConfig() (err error) {
	config := godotenv.Load("./pkg/config/.env")

	if config != nil {
		return nil
	}
	configuration := &Configuration{
		MongoDb: os.Getenv("MONGODB_URI"),
		Port:    os.Getenv("PORT"),
	}

	Config = configuration
	return
}

func GetConfig() *Configuration {
	return Config
}
