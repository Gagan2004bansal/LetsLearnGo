package config

import (
	"os"
	"strconv"
)

type Config struct {
	PORT         int
	DATABASE_URL string
}

func Load() (*Config, error) {
	cfg := &Config{
		PORT:         getEnvInt("PORT", 8082),
		DATABASE_URL: getEnvString("DATABASE_URL", ""),
	}

	return cfg, nil
}

func getEnvString(key string, default_value string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return default_value
}

func getEnvInt(key string, default_value int) int {
	if val, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return default_value
}
