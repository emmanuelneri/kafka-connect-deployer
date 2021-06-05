package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

const (
	DefaultUrl           = "http://localhost:8083"
	DefaultMaxRetry      = 3
	DefaultWaitStartTime = 1 * time.Nanosecond
)

type Config struct {
	ConnectUrl    string
	ConnectorsDir string
	MaxRetry      int
	WaitStartTime time.Duration
}

func New() (Config, error) {
	filesDir, err := connectorsDir()
	if err != nil {
		return Config{}, err
	}

	retry, err := maxRetry()
	if err != nil {
		return Config{}, err
	}

	startTime, err := waitStartTime()
	if err != nil {
		return Config{}, err
	}

	return Config{
		ConnectUrl:    kafkaConnectUrl(),
		ConnectorsDir: filesDir,
		MaxRetry:      retry,
		WaitStartTime: startTime,
	}, nil
}

func connectorsDir() (string, error) {
	envValue := os.Getenv("CONNECTORS_FILES_DIR")
	if envValue == "" {
		return "", errors.New("CONNECTORS_FILES_DIR empty")
	}

	return envValue, nil
}

func kafkaConnectUrl() string {
	envValue := os.Getenv("KAFKA_CONNECT_URL")
	if envValue == "" {
		return DefaultUrl
	}

	return envValue
}

func maxRetry() (int, error) {
	envValue := os.Getenv("MAX_RETRY")
	if envValue == "" {
		return DefaultMaxRetry, nil
	}

	return strconv.Atoi(envValue)
}

func waitStartTime() (time.Duration, error) {
	envValue := os.Getenv("WAIT_START_TIME")
	if envValue == "" {
		return 1 * time.Nanosecond, nil
	}

	return time.ParseDuration(envValue)
}
