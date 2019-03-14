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
	// Running app locally, need to pass these in from CLI
	// ImageDataDir    = "/Users/mr_trashcans/go/src/github.com/z0mi3ie/goimgs/data/www/images"
	// ImageServerHost = "http://localhost"

	ImageDataDir    = "/data/www/images"
	ImageServerHost = "http://imgserver"
	ImageServerPort = "8080"
	ImageServerPath = "/images/"
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
