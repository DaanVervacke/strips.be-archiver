package config

import (
	"fmt"
	"github.com/greencoda/confiq"
	confiqyaml "github.com/greencoda/confiq/loaders/yaml"

	"github.com/DaanVervacke/strips.be-archiver/internal/types"
)

type Config struct {
	API  types.API
	Auth types.Auth
}

func LoadConfig(configPath string) (Config, error) {
	configSet := confiq.New()

	var config Config

	if err := configSet.Load(
		confiqyaml.Load().FromFile(configPath),
	); err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	if err := configSet.Decode(&config, confiq.AsStrict()); err != nil {
		return Config{}, fmt.Errorf("failed to parse config fields: %w", err)
	}

	return config, nil
}
