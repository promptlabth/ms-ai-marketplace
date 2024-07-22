package config

import (
	"log/slog"
	"sync"
	"time"

	"github.com/caarlos0/env/v10"
)

type config struct {
	Port             string        `env:"PORT" envDefault:"8080"`
	AllowedOrigins   []string      `env:"ALLOWED_ORIGINS"`
	AllowedMethods   []string      `env:"ALLOWED_METHODS" envDefault:"GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS"`
	AllowedHeaders   []string      `env:"ALLOWED_HEADERS" envDefault:"X-Requested-With,Authorization,Origin,Content-Length,Content-Type"`
	AllowCredentials bool          `env:"ALLOW_CREDENTIALS" envDefault:"true"`
	CorsMaxAge       time.Duration `env:"CORS_MAX_AGE" envDefault:"5s"`
	GCP              GCP           `envPrefix:"GCP_"`
	NLPApiKey        ApiKey        `emvPrefix:"API_KEY_"`
	Database         Database      `envPrefix:"DB_"`
}

type ApiKey struct {
	Anthropic string `env:"ANTHROPIC"`
}
type Database struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
	SSLMode  string `env:"SSL"`
}

type GCP struct {
	GoogleAppleciationCredential string `env:"GOOGLE_APPLICATION_CREDENTIALS"`
	ProjectId                    string `env:"PROJECT_ID"`
}

var Val config
var once sync.Once

func init() {
	once.Do(
		func() {
			cfg := config{}
			if err := env.Parse(&cfg); err != nil {
				slog.Warn(err.Error())
			}

			Val = cfg
		},
	)
}
