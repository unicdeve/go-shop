package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config is the main configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
	SMTP     SMTPConfig
}

// ServerConfig is the configuration for the server
type ServerConfig struct {
	Port    string
	GinMode string
}

// DatabaseConfig is the configuration for the database
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig is the configuration for the JWT
type JWTConfig struct {
	Secret           string
	ExpiresIn        time.Duration
	RefreshExpiresIn time.Duration
}

// AWSConfig is the configuration for the AWS
type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string
	S3Region        string
	EventQueueName  string
}

// UploadConfig is the configuration for the upload
type UploadConfig struct {
	UploadProvider string
	Path           string
	MaxFileSize    int64
}

// SMTPConfig is the configuration for the SMTP server
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// LoadConfig loads the configuration from the environment variables
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	jwtRefreshExpiresIn, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRES_IN", "720h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "1024000"), 10, 64)
	smtpPort, _ := strconv.Atoi(getEnv("SMTP_PORT", "1025"))

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "go_shop"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:           getEnv("JWT_SECRET", "secret"),
			ExpiresIn:        jwtExpiresIn,
			RefreshExpiresIn: jwtRefreshExpiresIn,
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "us-east-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			S3Bucket:        getEnv("AWS_S3_BUCKET", "go_shop"),
			S3Endpoint:      getEnv("AWS_S3_ENDPOINT", "http://localhost:4566"),
			S3Region:        getEnv("AWS_S3_REGION", "us-east-1"),
			EventQueueName:  getEnv("AWS_EVENT_QUEUE_NAME", "ecommerce-events"),
		},
		Upload: UploadConfig{
			UploadProvider: getEnv("UPLOAD_PROVIDER", "local"),
			Path:           getEnv("UPLOAD_PATH", "uploads"),
			MaxFileSize:    maxUploadSize,
		},
		SMTP: SMTPConfig{
			Host:     getEnv("SMTP_HOST", "localhost"),
			Port:     smtpPort,
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("SMTP_FROM", "noreply@example.com"),
		},
	}, nil
}
