package config

import (
	"os"
)

// FilePathConfig -> Lokasi penyimpanan file upload atau untuk kebutuhan get file
type FilePathConfig struct {
	FilePath string
}

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	DbConfig
	FilePathConfig
}

func (c Config) readConfig() Config {
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),//set DB_HOST=localhost
		Port:     os.Getenv("DB_PORT"), //set DB_PORT=5432
		DbName:   os.Getenv("DB_NAME"), //set DB_USER=postgres
		User:     os.Getenv("DB_USER"), //set DB_PASSWORD=12345678
		Password: os.Getenv("DB_PASSWORD"), //set DB_NAME=db_gin_latihan
	}
	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"), //set API_HOST=localhost
		ApiPort: os.Getenv("API_PORT"), //set API_PORT=8888
	}
	c.FilePathConfig = FilePathConfig{
		FilePath: os.Getenv("FILE_PATH"), //set FILE_PATH=C:\imagesGolang
	}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	return cfg.readConfig()
}
