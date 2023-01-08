package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.poet", "5432")
}
