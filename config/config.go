package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppConfig struct {
}

type DatabaseConfig struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DbName       string
	TablePrefix  string
	Charset      string
	ParserTime   bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTConfig struct {
	Secret string
	Issuer string
	Expire time.Duration
}

var gConfig Config

func Load() error {
	vp := viper.New()
	vp.AddConfigPath("config/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	if err := vp.UnmarshalKey("Server", &gConfig.Server); err != nil {
		return err
	}
	gConfig.Server.ReadTimeout *= time.Second
	gConfig.Server.WriteTimeout *= time.Second
	// if err := vp.UnmarshalKey("App", &gConfig.App); err != nil {
	// 	return err
	// }
	// if err := vp.UnmarshalKey("Database", &gConfig.Database); err != nil {
	// 	return err
	// }
	if err := vp.UnmarshalKey("JWT", &gConfig.JWT); err != nil {
		return err
	}
	gConfig.JWT.Expire *= time.Second
	return nil
}

func GetConfig() *Config {
	return &gConfig
}
