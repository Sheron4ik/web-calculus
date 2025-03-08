package config

import "os"

var defaultCfg = map[string]string{
	"PORT":                    "8080",
	"TIME_ADDITION_MS":        "500",
	"TIME_SUBTRACTION_MS":     "500",
	"TIME_MULTIPLICATIONS_MS": "1000",
	"TIME_DIVISIONS_MS":       "1250",
	"COMPUTING_POWER":         "3",
}

type Config struct {
	Port           string
	TimeAddMs      string
	TimeSubMs      string
	TimeMulMs      string
	TimeDivMs      string
	ComputingPower string
}

func New() *Config {
	return &Config{
		Port:           getEnv("PORT"),
		TimeAddMs:      getEnv("TIME_ADDITION_MS"),
		TimeSubMs:      getEnv("TIME_SUBTRACTION_MS"),
		TimeMulMs:      getEnv("TIME_MULTIPLICATIONS_MS"),
		TimeDivMs:      getEnv("TIME_DIVISIONS_MS"),
		ComputingPower: getEnv("COMPUTING_POWER"),
	}
}

func getEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultCfg[key]
	}
	return env
}
