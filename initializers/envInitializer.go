package initializers

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvTest        Environment = "test"
	EnvStaging     Environment = "staging"
	EnvProduction  Environment = "production"
)

func (e *Environment) Decode(value string) error {
	value = strings.ToLower(value)
	switch Environment(value) {
	case EnvDevelopment, EnvTest, EnvStaging, EnvProduction:
		*e = Environment(value)
		return nil
	default:
		return fmt.Errorf("invalid ENV value: %s", value)
	}
}

type Config struct {
	Port      int         `envconfig:"PORT" default:"8080"`
	Env       Environment `envconfig:"ENV" default:"test"`
	DbUrl     string      `envconfig:"DB_URL" required:"true"`
	Debug     bool        `envconfig:"DEBUG" default:"false"`
	JwtSecret string      `envconfig:"JWT_SECRET"`
	JwtExpire string      `envconfig:"JWT_EXPIRE" default:"18h"`
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := envconfig.Process("", &AppConfig); err != nil {
		errorColor := color.New(color.FgRed, color.Bold).SprintFunc()
		log.Fatal(errorColor("Error: Failed to load environment variables → "), err)
	}
}

func LoadConfig() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return &cfg
}
