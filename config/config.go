package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Identity struct {
	Address string
}

type HoneyBadger struct {
	NetworkSize     int
	Byzantine       int
	BatchSize       int
	ProposeInterval time.Duration
}

type Members struct {
	Addresses []string
}

type Tpke struct {
	MasterPublicKey string
	Threshold       int
}

type Config struct {
	Identity    Identity
	HoneyBadger HoneyBadger
	Members     Members
	Tpke        Tpke
}

// TODO: change default config
var defaultConfig = &Config{
	Identity: Identity{
		Address: "localhost:15555",
	},
	HoneyBadger: HoneyBadger{
		NetworkSize:     10,
		Byzantine:       3,
		BatchSize:       3,
		ProposeInterval: 2 * time.Second,
	},
	Members: Members{
		Addresses: []string{
			"localhost:8080",
			"localhost:8081",
		},
	},
	Tpke: Tpke{
		MasterPublicKey: "asdf",
		Threshold:       3,
	},
}

var once sync.Once

var configPath = os.Getenv("HOME") + "/.cleisthenes/config.yml"

func Path() string {
	return configPath
}

func Get() *Config {
	once.Do(func() {
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			panic("cannot read config")
		}
		err := viper.Unmarshal(&defaultConfig)
		if err != nil {
			panic(fmt.Sprintf("error in read config, err: %s", err))
		}
	})
	return defaultConfig
}
