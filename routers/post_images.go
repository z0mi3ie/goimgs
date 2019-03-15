package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/config"
	"github.com/z0mi3ie/goimgs/db"
	"github.com/z0mi3ie/goimgs/image"
)

// UploadImages will upload given images metadata to the database for bookkeeping
// and then save the images to a directory, which can be served by the generated
// url for the image's ID in the database.
//
// Config values for targets, usernames, passwords etc. are pulled from the config
// package.
//
// NOTE: If an error is encountered here during the initial implementation the
// API will return a 500 -- not ideal and this will be cleaned up to make more sense :)
func UploadImages(c *gin.Context) {
	// Get the database client out of the context
	dbClient := c.MustGet(MySQLClientKey).(*db.Client)
	defer dbClient.DB().Close()

	// Get the config out of the context
	cfg := c.MustGet(ConfigKey).(*config.Config)

	fmt.Println("Uploading image")
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(500, err)
	}
	files := form.File["image"]

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
		res, err := stmt.Exec(imgData.ID(), imgData.URL(cfg.ImageServerTarget), imgData.OGName())
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
		targetDestination := fmt.Sprintf("%s/%s", cfg.ImageDataDir, imgData.Filename())
		if err := c.SaveUploadedFile(imgData.File(), targetDestination); err != nil {
			c.AbortWithError(500, err)
		}
	}
}
