package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
	Issuer string
	TTL    time.Duration
}

func Load() *Config {

	// Ignore error if .env doesn't exist (useful in production)
	_ = godotenv.Load()
	database := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Name:     getEnv("DB_NAME", "expense_tracker"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
	serverConfig := ServerConfig{
		Port: getEnv("SERVER_PORT", "8080"),
	}

	ttlHours, err := strconv.Atoi(os.Getenv("JWT_TTL_HOURS"))
	if err != nil {
		ttlHours = 24
	}
	jwtConfig := JWTConfig{
		Secret: getEnv("JWT_SECRET", ""),
		Issuer: getEnv("JWT_SECRET", ""),
		TTL:    time.Duration(ttlHours) * time.Hour,
	}
	config := &Config{
		Database: database,
		Server:   serverConfig,
		JWT:      jwtConfig,
	}

	config.validate()

	return config
}

func (c *Config) DatabaseURL() string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func (c *Config) validate() {

	if c.JWT.Secret == "" {
		log.Fatal("JWT_SECRET is required")
	}
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
