// Package config предоставляет функции для загрузки конфигурации приложения
package config

import "os"

// Config содержит конфигурационные параметры приложения
type Config struct {
	Port string
}

// Get возвращает конфигурацию, загруженную из переменных окружения
// Если переменная PORT не задана, используется значение по умолчанию "8080"
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
