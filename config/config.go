package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const configPath string = "./config/config.yml"

type (
	Config struct {
		HTTP     `yaml:"http"`
		Database `yaml:"database"`
	}

	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Database struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"DATABASE_POOL_MAX"`
		URL     string `env-required:"true" env:"DATABASE_URL"`
	}
)

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
