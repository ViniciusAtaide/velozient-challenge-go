package server

import "os"

type Config struct {
	BackendCors string
	AesKey      string
}

func ProvideConfig() *Config {
	return &Config{
		BackendCors: os.Getenv("ALLOWED_HOSTS"),
		AesKey:      os.Getenv("AES_KEY"),
	}
}
