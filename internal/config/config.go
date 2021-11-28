package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project Project `yaml:"project"`
	Grpc    Grpc    `yaml:"grpc"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) (Config, error) {
	var cfg Config

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return Config{}, err
	}
	defer func() { //nolint:gosec
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return cfg, nil
}
