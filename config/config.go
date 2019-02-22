package config

// Application Server configuration
const (
	ServerPort        = 8081
	ServerMaxFileSize = 1028 * 3
)

// Image Web Server Configuration
const (
	// The directory uploaded images are saved which are served from
	ImageDataDir    = "/Users/mr_trashcans/go/src/github.com/z0mi3ie/goimgs/data/"
	ImageServerHost = "localhost"
	ImageServerPort = "8082"
	ImageServerPath = "/data/"
)

// MySQL DB Configuration
const (
	MySQLUser     = "root"
	MySQLPassword = "password"
	MySQLDatabase = "mydb"
)
