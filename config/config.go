package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Stage string

const (
	Local Stage = "local"
	Dev         = "dev"
	Prod        = "prod"
)

type Config struct {
	ServerAddr string `envconfig:"ADDR"`
	Stage      Stage  `envconfig:"STAGE"`
	DBConfig
}

type DBConfig struct {
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	DBName     string `envconfig:"DB_NAME"`
	DBOptions  string `envconfig:"DB_OPTIONS"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
}

var cfgSync sync.Once
var confSingleton Config

func Get() *Config {
	cfgSync.Do(func() {
		err := envconfig.Process("", &confSingleton)

		if err != nil {
			panic(fmt.Sprintln("Couldn't process config", err))
		}
	})

	return &confSingleton
}

func (d *DBConfig) GetDBConnectionString() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s",
		d.DBHost, d.DBUser, d.DBPassword, d.DBName, d.DBPort, strings.Trim(d.DBOptions, "'"))
	return dsn
}
