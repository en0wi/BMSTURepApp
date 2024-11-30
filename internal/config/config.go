package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Environment      string `yaml:"env" env-default:"dev"`
	StoragePath      string `yaml:"storage_path" env-required:"true"`
	ConnectionString string `yaml:"connection_string" env-required:"true"`
	HTTPServer       `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"address"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func Load() Config {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Println("CONFIG_PATH is empty")
		configPath = "config/local.yaml"
	}

	// checks if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config file: %s", err)
	}

	return cfg
}
