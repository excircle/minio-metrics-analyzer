package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

/*
# Example config.yaml file
---
-
minio:                          # Config struct
  endpoint: 10.0.0.34           # MinioConfig struct
  accessKeyID: miniominio       # MinioConfig struct
  secretAccessKey: miniominio   # MinioConfig struct
  useSSL: false                 # MinioConfig struct

*/

type Config struct {
	Minio MinioConfig `yaml:"minio"`
}

// MinioConfig holds the configuration details for connecting to your MinIO server.
type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	UseSSL          bool   `yaml:"useSSL"`
}

// LoadConfig reads a YAML configuration file from the provided path and unmarshals it into the Config struct.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
