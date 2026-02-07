package configs

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load() (Config, error) {
	k := koanf.New(".")

	if err := k.Load(
		file.Provider("config.yaml"),

		yaml.Parser()); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
