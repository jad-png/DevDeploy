package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Docker   DockerConfig
}

type AppConfig struct {
	Environment string
	Port        int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host string
	Port int
	//Password string
}

type DockerConfig struct {
	Host string
}

func load() (*Config, error) {
	app, err := loadAppConfig()
	if err != nil {
		return nil, fmt.Errorf("load app config: %w", err)
	}

	db, err := loadDatabaseConfig()
	if err != nil {
		return nil, fmt.Errorf("load database config: %w", err)
	}

	redis, err := loadRedisConfig()
	if err != nil {
		return nil, fmt.Errorf("load redis config: %w", err)
	}

	docker, err := loadDockerConfig()
	if err != nil {
		return nil, fmt.Errorf("load docker config: %w", err)
	}

	return &Config{
		App:      app,
		Database: db,
		Redis:    redis,
		Docker:   docker,
	}, nil
}

func loadAppConfig() (AppConfig, error) {
	env, err := getEnv("APP_ENV")
	if err != nil {
		return AppConfig{}, err
	}

	port, err := getEnvAsInt("APP_PORT")
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Environment: env,
		Port:        port,
	}, nil
}

func loadDatabaseConfig() (DatabaseConfig, error) {
	host, err := getEnv("DB_HOST")
	if err != nil {
		return DatabaseConfig{}, err
	}

	port, err := getEnvAsInt("DB_PORT")
	if err != nil {
		return DatabaseConfig{}, err
	}

	user, err := getEnv("DB_USER")
	if err != nil {
		return DatabaseConfig{}, err
	}

	password, err := getEnv("DB_PASSWORD")
	if err != nil {
		return DatabaseConfig{}, err
	}

	name, err := getEnv("DB_NAME")
	if err != nil {
		return DatabaseConfig{}, err
	}

	return DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}, nil
}

func loadRedisConfig() (RedisConfig, error) {
	host, err := getEnv("REDIS_HOST")
	if err != nil {
		return RedisConfig{}, err
	}

	port, err := getEnvAsInt("REDIS_PORT")
	if err != nil {
		return RedisConfig{}, err
	}

	//password := os.Getenv("REDIS_PASSWORD")

	return RedisConfig{
		Host: host,
		Port: port,
		//Password: password,
	}, nil
}

func loadDockerConfig() (DockerConfig, error) {
	host, err := getEnv("DOCKER_HOST")
	if err != nil {
		return DockerConfig{}, err
	}

	return DockerConfig{
		Host: host,
	}, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s is required : ", key)
	}

	return value, nil
}

func getEnvAsInt(key string) (int, error) {
	value, err := getEnv(key)

	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(value)

	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", number)
	}

	return number, nil
}
