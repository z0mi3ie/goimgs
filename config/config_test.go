package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testServerPort             = "1010"
	testServerMaxFileSize      = "1024"
	testServerPortInt          = 1010
	testServerMaxFileSizeInt64 = int64(1024)

	testMySQLUser     = "testUser"
	testMySQLPassword = "testPassword"
	testMySQLDatabase = "testDatabase"

	testImageDataDir    = "/test/"
	testImageServerHost = "localhost"
	testImageServerPort = "2020"
	testImageServerPath = "/images"
)

var testImageServerTarget = buildImageServerTarget(testImageServerHost, testImageServerPort, testImageServerPath)

var testConfig = &Config{
	ServerPort:        testServerPort,
	ServerMaxFileSize: testServerMaxFileSize,
	MySQLUser:         testMySQLUser,
	MySQLPassword:     testMySQLPassword,
	MySQLDatabase:     testMySQLDatabase,
	ImageDataDir:      testImageDataDir,
	ImageServerHost:   testImageServerHost,
	ImageServerPort:   testImageServerPort,
	ImageServerPath:   testImageServerPath,
	ImageServerTarget: testImageServerTarget,
}

var testEnvVars = map[string]string{
	envAppServerPort:   testServerPort,
	envAppMaxFileSize:  testServerMaxFileSize,
	envMySQLDatabase:   testMySQLDatabase,
	envMySQLPassword:   testMySQLPassword,
	envMySQLUser:       testMySQLUser,
	envImageDataDir:    testImageDataDir,
	envImageServerHost: testImageServerHost,
	envImageServerPort: testImageServerPort,
	envImageServerPath: testImageServerPath,
}

func TestNewServerConfig(t *testing.T) {
	testCases := []struct {
		tc  string
		cfg *Config
		env map[string]string
	}{
		{
			tc:  "successfully read in args from the config",
			cfg: testConfig,
			env: testEnvVars,
		},
	}
	for _, test := range testCases {
		t.Run(test.tc, func(t *testing.T) {
			setEnvVars(test.env)
			nc, err := NewServerConfig()
			assert.Nil(t, err)
			assert.Equal(t, test.cfg, nc)
		})
	}
}

func TestServerPortInt(t *testing.T) {
	t.Run("Successfully get ServerPort as an int", func(t *testing.T) {
		setEnvVars(testEnvVars)
		nc, err := NewServerConfig()
		assert.Nil(t, err)
		p := nc.ServerPortInt()
		assert.Equal(t, testServerPortInt, p)
	})
}

func TestServerMaxFileSizeInt64(t *testing.T) {
	t.Run("Successfully get ServerMaxFileSize as an int64", func(t *testing.T) {
		setEnvVars(testEnvVars)
		nc, err := NewServerConfig()
		assert.Nil(t, err)
		p := nc.ServerMaxFileSizeInt64()
		assert.Equal(t, testServerMaxFileSizeInt64, p)
	})
}

func setEnvVars(e map[string]string) {
	for k, v := range e {
		_ = os.Setenv(k, v)
	}
}
