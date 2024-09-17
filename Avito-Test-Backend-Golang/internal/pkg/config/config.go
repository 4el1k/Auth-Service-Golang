package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// TODO : learn consulapi "github.com/hashicorp/consul/api"

type Config struct {
	HTTPServer `yaml:"httpServer"`
	AuthJwt    `yaml:"authJwt"`

	Database    `yaml:"database"`
	Enviroment  string `yaml:"ENVIROMENT" yaml-default:"prod"`
	LogFilePath string `yaml:"LOG_FILE_PATH" yaml-default:"logfile.log"`
}

type HTTPServer struct {
	Address           string        `yaml:"address" yaml-default:":8080"`
	Timeout           time.Duration `yaml:"timeout" yaml-default:"4s"`
	IdleTimeout       time.Duration `yaml:"idleTimeout" yaml-default:"60s"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout" yaml-defualt:"10s"`
}

type Database struct {
	DBName string `yaml:"DB_NAME" yaml-required:"true"`
	DBPass string `yaml:"DB_PASS" yaml-required:"true"`
	DBHost string `yaml:"DB_HOST" yaml-default:"0.0.0.0"`
	DBPort int    `yaml:"DB_PORT" yaml-required:"true"`
	DBUser string `yaml:"DB_USER" yaml-required:"true"`
}

type AuthJwt struct {
	AccessExpirationTime    time.Duration `yaml:"accessExpirationTime" yaml-defualt:"6h"`
	RefreshTokenTimeToLimit time.Duration `yaml:"RefreshTokenTimeToLimit" yaml-defualt:"6h"`
	Issuer                  string        `yaml:"issuer" yaml-defualt:"avito-banner"`
	Secret                  string        `yaml:"secret" yaml-defualt:"avito-banner"`
}

func MustLoad(filename string) *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Printf("cannot read .env file: %s\n (fix: you need to put .env file in main dir)", err)
		os.Exit(1)
	}

	// check if config file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Printf("config file does not exist: %s", filename)
		os.Exit(1)
	}

	if err := cleanenv.ReadConfig(filename, &cfg); err != nil {
		log.Printf("cannot read %s: %v", filename, err)
		os.Exit(1)
	}

	return &cfg
}
