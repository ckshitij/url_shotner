package config

import "os"

var Config *ServiceConfig

type ServiceConfig struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host         string
	Port         string
	ReadTimeout  int // ReadTimeout values in Seconds
	WriteTimeout int // WriteTimeout values in Seconds
	IdleTimeout  int // IdleTimeout values in Seconds
}

func LoadServiceConfig() *ServiceConfig {
	Config = &ServiceConfig{
		Server: ServerConfig{
			Host:         getEnv("SERVICE_HOST", "localhost"),
			Port:         getEnv("SERVICE_PORT", "8088"),
			ReadTimeout:  10,
			WriteTimeout: 10,
			IdleTimeout:  60,
		},
	}
	return Config
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
