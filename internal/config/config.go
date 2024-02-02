package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

const (
	configFilePath = "config/local.yaml"
)

type Config struct {
	Env         string        `yaml:"env" env-requited:"true"`
	StoragePath string        `yaml:"storage_path" env-requited:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-requited:"true"`
	Grpc        GRPCConfig    `yaml:"grpc" env-requited:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-requited:"true"`
	Timeout time.Duration `yaml:"timeout" env-requited:"true"`
}

func MustLoad() *Config {
	c := new(Config)

	f, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}

	if err := cleanenv.ParseYAML(f, c); err != nil {
		panic(err)
	}

	return c
}
