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
	ServerPort        int
	ServerMaxFileSize int64

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

// NewServerConfig generates and returns an object with all of the configuration values
func NewServerConfig() (*Config, error) {
	c := &Config{}

	p := utils.MustGetEnvironmentVariable(envAppServerPort)
	i, err := strconv.Atoi(p)
	if err != nil {
		return nil, err
	}
	c.ServerPort = i
	s := utils.MustGetEnvironmentVariable(envAppMaxFileSize)
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	c.ServerMaxFileSize = i64

	c.ImageDataDir = utils.MustGetEnvironmentVariable(envImageDataDir)
	c.ImageServerHost = utils.MustGetEnvironmentVariable(envImageServerHost)
	c.ImageServerPort = utils.MustGetEnvironmentVariable(envImageServerPort)
	c.ImageServerPath = utils.MustGetEnvironmentVariable(envImageServerPath)
	c.ImageServerTarget = fmt.Sprintf("%s:%s%s", c.ImageServerHost, c.ImageServerPort, c.ImageServerPath)

	c.MySQLDatabase = utils.MustGetEnvironmentVariable(envMySQLDatabase)
	c.MySQLUser = utils.MustGetEnvironmentVariable(envMySQLUser)
	c.MySQLPassword = utils.MustGetEnvironmentVariable(envMySQLPassword)

	return c, nil
}
