package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	DBMaxConnections string

	RedisURL      string
	RedisPassword string

	JWTSecret       string
	APIKey          string
	StripeSecretKey string
	SendGridAPIKey  string

	GoogleClientID     string
	GoogleClientSecret string
	GitHubClientID     string
	GitHubClientSecret string

	WebhookURL             string
	NotificationServiceURL string

	Environment string
	Debug       string
	LogLevel    string

	EncryptionKey string
	SigningKey    string
}

func LoadSOPSEnv(filename string) (*EnvConfig, error) {

	cmd := exec.Command("sops", "-d", filename)
	decryptedData, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt env file: %w", err)
	}

	tmpFile, err := os.CreateTemp("", "decrypted-*.env")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(decryptedData); err != nil {
		return nil, fmt.Errorf("failed to write temp file: %w", err)
	}
	tmpFile.Close()

	envMap, err := godotenv.Read(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to parse env file: %w", err)
	}

	config := &EnvConfig{
		DBHost:                 envMap["DB_HOST"],
		DBPort:                 envMap["DB_PORT"],
		DBName:                 envMap["DB_NAME"],
		DBUser:                 envMap["DB_USER"],
		DBPassword:             envMap["DB_PASSWORD"],
		DBMaxConnections:       envMap["DB_MAX_CONNECTIONS"],
		RedisURL:               envMap["REDIS_URL"],
		RedisPassword:          envMap["REDIS_PASSWORD"],
		JWTSecret:              envMap["JWT_SECRET"],
		APIKey:                 envMap["API_KEY"],
		StripeSecretKey:        envMap["STRIPE_SECRET_KEY"],
		SendGridAPIKey:         envMap["SENDGRID_API_KEY"],
		GoogleClientID:         envMap["GOOGLE_CLIENT_ID"],
		GoogleClientSecret:     envMap["GOOGLE_CLIENT_SECRET"],
		GitHubClientID:         envMap["GITHUB_CLIENT_ID"],
		GitHubClientSecret:     envMap["GITHUB_CLIENT_SECRET"],
		WebhookURL:             envMap["WEBHOOK_URL"],
		NotificationServiceURL: envMap["NOTIFICATION_SERVICE_URL"],
		Environment:            envMap["ENVIRONMENT"],
		Debug:                  envMap["DEBUG"],
		LogLevel:               envMap["LOG_LEVEL"],
		EncryptionKey:          envMap["ENCRYPTION_KEY"],
		SigningKey:             envMap["SIGNING_KEY"],
	}

	return config, nil
}

func LoadSOPSEnvToSystem(filename string) error {

	cmd := exec.Command("sops", "-d", filename)
	decryptedData, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to decrypt env file: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(decryptedData)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("failed to set env var %s: %w", key, err)
		}
	}

	return scanner.Err()
}

func PrintConfig(config *EnvConfig) {
	fmt.Println("🔓 Successfully loaded and decrypted environment configuration:")
	fmt.Println("================================================================")

	fmt.Println("\n📊 Database Configuration:")
	fmt.Printf("  DB_HOST: %s\n", config.DBHost)
	fmt.Printf("  DB_PORT: %s\n", config.DBPort)
	fmt.Printf("  DB_NAME: %s\n", config.DBName)
	fmt.Printf("  DB_USER: %s\n", config.DBUser)
	fmt.Printf("  DB_PASSWORD: %s\n", maskSecret(config.DBPassword))
	fmt.Printf("  DB_MAX_CONNECTIONS: %s\n", config.DBMaxConnections)

	fmt.Println("\n🔴 Redis Configuration:")
	fmt.Printf("  REDIS_URL: %s\n", maskSecret(config.RedisURL))
	fmt.Printf("  REDIS_PASSWORD: %s\n", maskSecret(config.RedisPassword))

	fmt.Println("\n🔐 API Keys & Secrets:")
	fmt.Printf("  JWT_SECRET: %s\n", maskSecret(config.JWTSecret))
	fmt.Printf("  API_KEY: %s\n", maskSecret(config.APIKey))
	fmt.Printf("  STRIPE_SECRET_KEY: %s\n", maskSecret(config.StripeSecretKey))
	fmt.Printf("  SENDGRID_API_KEY: %s\n", maskSecret(config.SendGridAPIKey))

	fmt.Println("\n🔑 OAuth Credentials:")
	fmt.Printf("  GOOGLE_CLIENT_ID: %s\n", config.GoogleClientID)
	fmt.Printf("  GOOGLE_CLIENT_SECRET: %s\n", maskSecret(config.GoogleClientSecret))
	fmt.Printf("  GITHUB_CLIENT_ID: %s\n", config.GitHubClientID)
	fmt.Printf("  GITHUB_CLIENT_SECRET: %s\n", maskSecret(config.GitHubClientSecret))

	fmt.Println("\n🌐 External Services:")
	fmt.Printf("  WEBHOOK_URL: %s\n", config.WebhookURL)
	fmt.Printf("  NOTIFICATION_SERVICE_URL: %s\n", config.NotificationServiceURL)

	fmt.Println("\n⚙️ Environment Settings:")
	fmt.Printf("  ENVIRONMENT: %s\n", config.Environment)
	fmt.Printf("  DEBUG: %s\n", config.Debug)
	fmt.Printf("  LOG_LEVEL: %s\n", config.LogLevel)

	fmt.Println("\n🔒 Encryption Keys:")
	fmt.Printf("  ENCRYPTION_KEY: %s\n", maskSecret(config.EncryptionKey))
	fmt.Printf("  SIGNING_KEY: %s\n", maskSecret(config.SigningKey))
}

func PrintSystemEnvVars() {
	fmt.Println("\n🌍 Environment Variables (loaded into system):")
	fmt.Println("================================================")

	envVars := os.Environ()
	sort.Strings(envVars)

	ourVars := []string{
		"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_MAX_CONNECTIONS",
		"REDIS_URL", "REDIS_PASSWORD",
		"JWT_SECRET", "API_KEY", "STRIPE_SECRET_KEY", "SENDGRID_API_KEY",
		"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "GITHUB_CLIENT_ID", "GITHUB_CLIENT_SECRET",
		"WEBHOOK_URL", "NOTIFICATION_SERVICE_URL",
		"ENVIRONMENT", "DEBUG", "LOG_LEVEL",
		"ENCRYPTION_KEY", "SIGNING_KEY",
	}

	for _, varName := range ourVars {
		if value := os.Getenv(varName); value != "" {
			if isSecret(varName) {
				fmt.Printf("  %s=%s\n", varName, maskSecret(value))
			} else {
				fmt.Printf("  %s=%s\n", varName, value)
			}
		}
	}
}

func maskSecret(value string) string {
	if len(value) <= 4 {
		return strings.Repeat("*", len(value))
	}
	return value[:2] + strings.Repeat("*", len(value)-4) + value[len(value)-2:]
}

func isSecret(varName string) bool {
	secretVars := []string{
		"PASSWORD", "SECRET", "KEY", "TOKEN", "CREDENTIAL", "PRIVATE",
	}

	upper := strings.ToUpper(varName)
	for _, secret := range secretVars {
		if strings.Contains(upper, secret) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("🔐 SOPS Environment Variable Manager")
	fmt.Println("=====================================")

	fmt.Println("\n📋 Method 1: Loading into structured configuration")
	config, err := LoadSOPSEnv("config.sops.env")
	if err != nil {
		log.Fatalf("Error loading SOPS env config: %v", err)
	}
	PrintConfig(config)

	fmt.Println("\n" + strings.Repeat("=", 60))

	fmt.Println("\n📋 Method 2: Loading into system environment variables")
	if err := LoadSOPSEnvToSystem("config.sops.env"); err != nil {
		log.Fatalf("Error loading SOPS env to system: %v", err)
	}
	PrintSystemEnvVars()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🚀 Example Usage:")

	fmt.Println("\n1️⃣ Using Structured Config:")
	fmt.Printf("   Database DSN: postgresql://%s:%s@%s:%s/%s\n",
		config.DBUser,
		maskSecret(config.DBPassword),
		config.DBHost,
		config.DBPort,
		config.DBName)

	fmt.Println("\n2️⃣ Using System Environment Variables:")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Printf("   Database DSN: postgresql://%s:%s@%s:%s/%s\n",
		dbUser,
		maskSecret(dbPassword),
		dbHost,
		dbPort,
		dbName)

	fmt.Println("\n✅ SOPS environment variable integration complete!")
	fmt.Println("Your environment secrets are now loaded and ready to use! 🎉")
}
