package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database	Connection
}

type Connection struct {
	Path 		string
	DatabaseURL string
}

//=================================================================================================================

// * Init Config
func InitConfig(env string) *Config {
	configPath := GetConfigPath(env)

	confFile, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", confFile)
	}

	conf, err := ParseConfig(confFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return conf
}

// * Mendapatkan Path Konfigurasi Local
func GetConfigPath(configPath string) string {
	if configPath == "test" {
		return "config-test"
	}
	return "./config/config-local"
}

// * Membaca informasi dari file konfigurasi
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return v, errors.New("config file not found")
		}
		return v, err
	}

	return v, nil
}

// * Parse dari file konfigurasi yaml ke struct
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
