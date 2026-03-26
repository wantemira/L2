package config

import "os"

type Config struct {
	Port string
}

func Get() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(k, def string) string {
	val := os.Getenv(k)
	if val == "" {
		return def
	}
	return val
}
