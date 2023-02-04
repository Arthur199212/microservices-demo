package utils

import "os"

type Config struct {
	AllowTestCardNumbers bool
	Port                 string
}

const (
	defaultPort = "5003"
)

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	allowTestCardNumbers := false
	if os.Getenv("ALLOW_TEST_CARD_NUMBERS") == "true" {
		allowTestCardNumbers = true
	}

	return &Config{
		AllowTestCardNumbers: allowTestCardNumbers,
		Port:                 port,
	}
}
