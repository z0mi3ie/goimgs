package main

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// Configuration for the server that we are hardcoding but will be changing later
// Most of this should come in as command line or environment config
const (
	port        = 8081
	datadir     = "/Users/mr_trashcans/go/src/github.com/z0mi3ie/goimgs/data/"
	maxFileSize = 1028 * 3
)

// Server is the instance of the server running and also holds the server wide
// configuration and fields
type Server struct {
	router *gin.Engine
}

// ImageData holds all associated data the DB needs to store for an image
// before saving that image to the file system
type ImageData struct {
	id     string
	url    string
	ogName string
}

// Health is a simple health check for the server
func Health(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

// generateImageID generates a new UUID and returns the string to be used for the
// image file that will be saved in the database.
func generateImageID() (string, error) {
	gid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return gid.String(), nil
}

// generateImageURL takes an image ID and generates the accessible
// URL for the image to be saved in the database.
func generateImageURL(id string, originalName string) (string, error) {
	// This is a real quick and dirty way to save the proper file type, should be fixed
	// and images should be normalized to all the same type, but we are going to use a
	// an image processing library for that.
	fileType := originalName[len(originalName)-3]
	// This will need to be updated to dynamically handle the endpoint
	target := fmt.Sprintf("localhost:8081/data/%s.%s", id, fileType)
	// This needs a lot more work to do the actual things and error handling
	return target, nil
}

func getDB(driver string, datasource string) *sql.DB {
	// For now, real quick and dirty, connect to the DB for this request
	db, err := sql.Open("mysql", "root:password@/mydb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	//defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}

// UploadImage will take the file from the request and store it in the /data directory (for now)
func UploadImage(c *gin.Context) {
	// Parse the multipart form into a Form
	form, _ := c.MultipartForm()
	// Grab the image files from the parsed form
	files := form.File["image"]
	fmt.Println(fmt.Sprintf("files %v", files))

	db := getDB("mysql", "root:password@/mydb")
	defer db.Close()

	for _, file := range files {
		printFileInfo(file)

		imgID, err := generateImageID()
		if err != nil {
			fmt.Println("Error generating UUID, this shouldn't happen, 500")
			c.AbortWithStatus(500)
		}

		imgURL, err := generateImageURL(imgID, file.Filename)
		if err != nil {
			fmt.Println("Error generating image url, this shouldn't happen, 500")
			c.AbortWithStatus(500)
		}

		imgData := &ImageData{
			id:     imgID,
			url:    imgURL,
			ogName: file.Filename,
		}

		stmt, err := db.Prepare("INSERT INTO image(id, url, og_name) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(imgData.id, imgData.url, imgData.ogName)
		if err != nil {
			log.Fatal(err)
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, targetFileDestination(datadir, file))
		if err != nil {
			fmt.Println("[ERROR] could not save image", err)
		}
	}
}

// GetImage retrieves the image from the filesystem
func GetImage(c *gin.Context) {
}

// GetImages retrieves all images
func GetImages(c *gin.Context) {
	db := getDB("mysql", "root:password@/mydb")
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM image")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rowImage := &ImageData{}
	for rows.Next() {
		err := rows.Scan(&rowImage.id, &rowImage.url, &rowImage.ogName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rowImage.id, rowImage.url, rowImage.ogName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func targetFileDestination(dir string, f *multipart.FileHeader) string {
	return fmt.Sprintf("%s/%s", dir, f.Filename)
}

func printFileInfo(f *multipart.FileHeader) {
	fmt.Println(fmt.Sprintf("Filename: %v", f.Filename))
	fmt.Println(fmt.Sprintf("Header: %v", f.Header))
	fmt.Println(fmt.Sprintf("Size: %v", f.Size))
}

func main() {

	server := &Server{
		router: gin.Default(),
	}

	db := getDB("mysql", "root:password@/mydb")
	db.Ping()

	// Setup config
	server.router.MaxMultipartMemory = maxFileSize

	// Setup Routes
	server.router.GET("/ping", Health)
	server.router.POST("/image", UploadImage)
	server.router.GET("/image", GetImage)
	server.router.GET("/images", GetImages)

	// Start Server
	server.router.Run(fmt.Sprintf(":%v", port)) // listen and serve on 0.0.0.0:8080
}
