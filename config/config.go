package config

import "fmt"

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

// ImageServerTarget returns the currently defined target directory where images
// are being served from
func ImageServerTarget() string {
	return fmt.Sprintf("%s:%s%s", ImageServerHost, ImageServerPort, ImageServerPath)
}
