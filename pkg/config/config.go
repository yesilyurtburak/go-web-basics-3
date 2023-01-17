package config

import "log"

// This package stores system wide configurations.

type AppConfig struct {
	InfoLog *log.Logger
}
