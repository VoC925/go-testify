package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/VoC925/go-testify/internal"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server  `yaml:"server"`
	Storage `yaml:"storage"`
}

type Server struct {
	Port string `yaml:"port" env-default:"8080"`
	Host string `yaml:"host" env-default:"localhost"`
}

type Storage struct {
	Name string `yaml:"database"`
}

var (
	once           sync.Once
	instanceConfig *Config
)

func New(path string) *Config {
	if path == "" {
		log.Fatal(internal.ErrEmptyConfigPath)
	}
	once.Do(func() {
		instanceConfig = &Config{}
		if err := cleanenv.ReadConfig(path, instanceConfig); err != nil {
			help, _ := cleanenv.GetDescription(instanceConfig, nil)
			log.Fatalf(fmt.Sprintf(`%s : %s`, err, help))
		}
	})
	return instanceConfig
}
