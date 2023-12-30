package config

import (
	"os"
)

type EncryptConfig struct {
	Key string
	IV  string
}

func LoadEncryptConfig() EncryptConfig {
	return EncryptConfig{
		Key: os.Getenv("ENCRYPT_KEY"),
		IV:  os.Getenv("ENCRYPT_IV"),
	}
}
