package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/db"
	"github.com/z0mi3ie/goimgs/image"
)

// GetImages retrieves all images
func GetImages(c *gin.Context) {
	dbClient := c.MustGet(MySQLClientKey).(*db.Client)
	defer dbClient.DB().Close()

	stmt, err := dbClient.DB().Prepare("SELECT * FROM image")
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer rows.Close()

	retrievedImgs := []image.MetaData{}
	for rows.Next() {
		var imgID, imgURL, imgOGName string
		var imgDeleted bool
		var imgDescription, imgTitle string

		err := rows.Scan(&imgID, &imgURL, &imgOGName, &imgDeleted, &imgDescription, &imgTitle)
		if err != nil {
			c.AbortWithError(500, err)
		}
		fmt.Println(imgID, imgURL, imgOGName, imgDescription, imgTitle)
		newImg := image.NewImageMetaData(imgID, imgURL, imgOGName, imgDescription, imgTitle)
		retrievedImgs = append(retrievedImgs, newImg)
	}
	if err = rows.Err(); err != nil {
		c.AbortWithError(500, err)
	}

	fmt.Println("Number of retrieved images: ", len(retrievedImgs))
	fmt.Println("The retrieved images names...")
	for _, img := range retrievedImgs {
		fmt.Println(">> ", img.ID)
	}

	c.JSON(200, retrievedImgs)
}
