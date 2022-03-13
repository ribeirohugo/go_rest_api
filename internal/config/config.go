package config

import (
	"fmt"
	"os"
)

const (
	serverHostEnv = "HOST"
	serverPortEnv = "PORT"

	databaseAddressEnv = "DATABASE"
)

type Database struct {
	Address string
}

type Server struct {
	Host string
	Port string
}

type Config struct {
	Database Database
	Server   Server
}

func Load() Config {
	return Config{
		Database: Database{
			Address: os.Getenv(databaseAddressEnv),
		},
		Server: Server{
			Host: os.Getenv(serverHostEnv),
			Port: os.Getenv(serverPortEnv),
		},
	}
}

func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}
