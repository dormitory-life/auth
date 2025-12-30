package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const dbConnectionStringTemplate = "%s://%s:%s@%s:%d/%s?sslmode=%s"

type Config struct {
	Env    string         `yaml:"env"`
	Db     DataBaseConfig `yaml:"database"`
	Server ServerConfig   `yaml:"server"`
}

type DataBaseConfig struct {
	Driver         string `yaml:"driver"`
	Host           string `yaml:"host"`
	Port           uint16 `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Database       string `yaml:"database"`
	SSLMode        string `yaml:"ssl_mode"`
	MigrationsPath string `yaml:"migrations_path"`
}

type ServerConfig struct {
	Port uint16 `yaml:"port"`
}

func ParseConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return config, nil
}

func (c *DataBaseConfig) GetConnectionString() string {
	return fmt.Sprintf(
		dbConnectionStringTemplate,
		c.Driver,
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.SSLMode,
	)
}
