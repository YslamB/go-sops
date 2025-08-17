package main

import (
	"fmt"
	"log"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Storage Storage `yaml:"storage"`
	JWT     JWT     `yaml:"jwt"`
}

type Storage struct {
	PSQL  PSQL  `yaml:"psql"`
	Redis Redis `yaml:"redis"`
}

type PSQL struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Database      string `yaml:"database"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	PGPoolMaxConn int    `yaml:"pg_pool_max_conn"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWT struct {
	Auth string `yaml:"auth"`
}

func LoadSOPSConfig(filename string) (*Config, error) {
	cmd := exec.Command("sops", "-d", filename)
	decryptedData, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(decryptedData, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}

func main() {
	config, err := LoadSOPSConfig("config.sops.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Println("🔓 Successfully loaded and decrypted configuration:")
	fmt.Println("====================================================")

	fmt.Println("📊 Database Configuration:")
	fmt.Printf("  Host: %s\n", config.Storage.PSQL.Host)
	fmt.Printf("  Port: %d\n", config.Storage.PSQL.Port)
	fmt.Printf("  Database: %s\n", config.Storage.PSQL.Database)
	fmt.Printf("  Username: %s\n", config.Storage.PSQL.Username)
	fmt.Printf("  Password: %s\n", config.Storage.PSQL.Password)
	fmt.Printf("  Max Connections: %d\n", config.Storage.PSQL.PGPoolMaxConn)

	fmt.Println("\n🔴 Redis Configuration:")
	fmt.Printf("  Address: %s\n", config.Storage.Redis.Addr)
	fmt.Printf("  Port: %d\n", config.Storage.Redis.Port)
	fmt.Printf("  Username: %s\n", config.Storage.Redis.Username)
	fmt.Printf("  Password: %s\n", config.Storage.Redis.Password)
	fmt.Printf("  Database: %d\n", config.Storage.Redis.DB)

	fmt.Println("\n🔐 JWT Configuration:")
	fmt.Printf("  Auth Key: %s\n", config.JWT.Auth)

	fmt.Println("\n======================================================")
	fmt.Println("🚀 Example Usage:")
	fmt.Printf("PostgreSQL DSN: postgresql://%s:%s@%s:%d/%s\n",
		config.Storage.PSQL.Username,
		config.Storage.PSQL.Password,
		config.Storage.PSQL.Host,
		config.Storage.PSQL.Port,
		config.Storage.PSQL.Database)
	fmt.Printf("Redis URL: redis://%s:%s@%s:%d/%d\n",
		config.Storage.Redis.Username,
		config.Storage.Redis.Password,
		config.Storage.Redis.Addr,
		config.Storage.Redis.Port,
		config.Storage.Redis.DB)
}
