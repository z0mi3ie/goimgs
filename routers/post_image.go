package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/config"
	"github.com/z0mi3ie/goimgs/db"
	"github.com/z0mi3ie/goimgs/image"
)

// UploadImage will upload given images metadata to the database for bookkeeping
// and then save the images to a directory, which can be served by the generated
// url for the image's ID in the database.
//
// Config values for targets, usernames, passwords etc. are pulled from the config
// package.
//
// NOTE: If an error is encountered here during the initial implementation the
// API will return a 500 -- not ideal and this will be cleaned up to make more sense :)
func UploadImage(c *gin.Context) {
	fmt.Println("Uploading image")
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(500, err)
	}
	files := form.File["image"]
	dbClient, err := db.NewMySQLClient(config.MySQLUser, config.MySQLPassword, config.MySQLDatabase)
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer dbClient.DB().Close()
	for _, file := range files {
		// Add file data to the database
		imgData, err := image.NewImageData(file)
		if err != nil {
			c.AbortWithError(500, err)
		}
		stmt, err := dbClient.DB().Prepare("INSERT INTO image(id, url, og_name) VALUES(?,?,?)")
		if err != nil {
			c.AbortWithError(500, err)
		}
		res, err := stmt.Exec(imgData.ID(), imgData.URL(config.ImageServerTarget()), imgData.Filename())
		if err != nil {
			c.AbortWithError(500, err)
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			c.AbortWithError(500, err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			c.AbortWithError(500, err)
		}
		fmt.Printf("ID = %d, affected = %d\n", lastID, rowCnt)

		// Save the file to the target destination
		targetDestination := fmt.Sprintf("%s/%s", config.ImageDataDir, imgData.Filename())
		if err := c.SaveUploadedFile(imgData.File(), targetDestination); err != nil {
			c.AbortWithError(500, err)
		}
	}
}
