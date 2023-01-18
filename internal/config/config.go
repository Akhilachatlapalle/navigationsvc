package config

import (
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const (
	DefaultPort     = "5055"
	ShutdownTimeout = 20 * time.Second
)

type Config struct {
	SectorID int

	Port string

	Logger *zap.Logger
}

func New() (*Config, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sectorID := 1
	sectorIDString := os.Getenv("SECTOR_ID")
	if sectorIDString != "" {
		sectorID, err = strconv.Atoi(sectorIDString)
		if err != nil {
			return nil, err
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	result := Config{
		Logger: logger,

		Port: port,

		SectorID: sectorID,
	}
	return &result, nil
}
