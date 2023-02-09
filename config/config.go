package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	TgToken       string        `yaml:"tg_token"`
	StorageConfig StorageConfig `yaml:"storage_config"`
	Buttons       Buttons       `yaml:"buttons"`
}

type StorageConfig struct {
	Host           string        `json:"host"`
	Port           string        `json:"port"`
	Database       string        `json:"database"`
	Username       string        `json:"username"`
	Password       string        `json:"password"`
	MaxAttempts    int           `json:"maxattempts"`
	MaxDelaySecond time.Duration `json:"maxdelaysecond"`
}

type Buttons struct {
	BtStart string `json:"btstart"`
	BtNext  string `json:"btnext"`
}

//GetConfig unmarshal and return config file fields
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			description, _ := cleanenv.GetDescription(instance, nil)
			log.Errorf("%v, error description:%v", description, err)
		}
	})
	return instance
}
