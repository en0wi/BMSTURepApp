package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Environment string `yaml:"env" env-default:"dev"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"address"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func Load() Config {

	defaultConfigPath := "config/local.yaml"
	if err := os.Setenv("CONFIG_PATH", defaultConfigPath); err != nil {
		fmt.Println("Error setting environment variable:", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is empty")
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
