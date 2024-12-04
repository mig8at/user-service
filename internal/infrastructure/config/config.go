package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	BadgerPath string
	SqlitePath string
	Env        string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer la configuración: %v", err)
	}

	return &Config{
		Port:       viper.GetString("server.port"),
		BadgerPath: viper.GetString("db.badger"),
		SqlitePath: viper.GetString("db.sqlite"),
		Env:        viper.GetString("env"),
	}
}
