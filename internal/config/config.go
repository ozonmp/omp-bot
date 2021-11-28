package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type cnmFilm struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type logger struct {
	LogLevel string `yaml:"log_level"`
}

type Config struct {
	CnmFilm cnmFilm `yaml:"cnm_film"`
	Logger logger `yaml:"logger"`
}

func InitConfigYAML(filePath string) (*Config, error) {
	cfg := &Config{}

	file, err := os.Open(filePath)
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	} ()
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
