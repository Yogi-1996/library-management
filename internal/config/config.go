package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string `yaml:"Port" env:"APP_PORT"`
	Env  string `yaml:"Env" env:"APP_ENV"`
}

type DBConfig struct {
	Host     string `yaml:"Host" env:"DB_HOST" env-required:"true"`
	Port     int    `yaml:"Port" env:"DB_PORT" env-required:"true"`
	User     string `yaml:"User" env:"DB_USER" env-required:"true"`
	Password string `yaml:"Password" env:"DB_PASSWORD" env-required:"true"`
	Name     string `yaml:"Name" env:"DB_NAME" env-required:"true"`
	SSLMode  string `yaml:"SSLMode" env:"DB_SSLMODE" env-required:"true"`
}

type SecurityConfig struct {
	Key   string `yaml:"Key" env:"SEC_KEY" env-default:"librarymanagment"`
	Hours int    `yaml:"Hours" env:"SEC_HOURS" env-default:"24"`
	Hash  int    `yaml:"Hash" env:"SEC_HASH" env-default:"10"`
}

type Config struct {
	App AppConfig      `yaml:"App"`
	DB  DBConfig       `yaml:"DB"`
	Sec SecurityConfig `yaml:"Sec"`
}

func LoadConfig() *Config {
	var env string

	flags := flag.String("env", "", "enviorment variable")
	flag.Parse()

	_ = godotenv.Load(".ENV")
	env = os.Getenv("APP_ENV")

	if *flags != "" {
		env = *flags
	} else {
		env = env
	}

	if env == "" {
		env = "dev"
	}

	configPath := fmt.Sprintf("config/%s.yaml", env)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can not read config : %s", err)
	}

	return &cfg

}
