package config

import (
	"fmt"
	"strconv"

	"github.com/z0mi3ie/goimgs/utils"
)

const (
	envMySQLDatabase   = "MYSQL_DATABASE"
	envMySQLUser       = "MYSQL_USER"
	envMySQLPassword   = "MYSQL_PASSWORD"
	envAppServerPort   = "APP_SERVER_PORT"
	envAppMaxFileSize  = "APP_MAX_FILE_SIZE"
	envImageDataDir    = "IMAGE_DATA_DIR"
	envImageServerHost = "IMAGE_SERVER_HOST"
	envImageServerPort = "IMAGE_SERVER_PORT"
	envImageServerPath = "IMAGE_SERVER_PATH"
)

// Config holds all of the configuration for the app
type Config struct {
	// API Config
	ServerPort        string
	ServerMaxFileSize string

	// MySQL DB Configuration
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	// NGINX Image Server host and directory information
	ImageDataDir      string
	ImageServerHost   string
	ImageServerPort   string
	ImageServerPath   string
	ImageServerTarget string
}

// NewServerConfig generates and returns a struct to hold all of the service's configuration values
func NewServerConfig() (*Config, error) {
	c := &Config{}

	c.ServerPort = utils.MustGetEnvironmentVariable(envAppServerPort)
	c.ServerMaxFileSize = utils.MustGetEnvironmentVariable(envAppMaxFileSize)

	c.ImageDataDir = utils.MustGetEnvironmentVariable(envImageDataDir)
	c.ImageServerHost = utils.MustGetEnvironmentVariable(envImageServerHost)
	c.ImageServerPort = utils.MustGetEnvironmentVariable(envImageServerPort)
	c.ImageServerPath = utils.MustGetEnvironmentVariable(envImageServerPath)
	c.ImageServerTarget = buildImageServerTarget(c.ImageServerHost, c.ImageServerPort, c.ImageServerPath)

	c.MySQLDatabase = utils.MustGetEnvironmentVariable(envMySQLDatabase)
	c.MySQLUser = utils.MustGetEnvironmentVariable(envMySQLUser)
	c.MySQLPassword = utils.MustGetEnvironmentVariable(envMySQLPassword)

	return c, nil
}

func buildImageServerTarget(host string, port string, path string) string {
	return fmt.Sprintf("%s:%s%s", host, port, path)
}

// ServerPortInt returns the server port the config has read in as an
// int for use where int is needed, but it panics if there is an
// error encountered since we need the server port to be valid to
// start the server
func (c *Config) ServerPortInt() int {
	i, err := strconv.Atoi(c.ServerPort)
	if err != nil {
		panic(fmt.Sprintf("ServerPort is not parseable to int: %s", c.ServerPort))
	}
	return i
}

// ServerMaxFileSizeInt64 returns the max file size the config has read in as an
// int64 for use where int64 is needed. A panic occurs if there is an error
// encountered since we need a valid max file size to start the server
func (c *Config) ServerMaxFileSizeInt64() int64 {
	i64, err := strconv.ParseInt(c.ServerMaxFileSize, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("ServerMaxFileSize is not parseable to int: %s", c.ServerMaxFileSize))
	}
	return i64
}
