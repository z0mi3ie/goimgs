package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/db"
)

// DeleteImages removes desired images through the query parameter "id"
func DeleteImages(c *gin.Context) {
	dbClient := c.MustGet(MySQLClientKey).(*db.Client)
	defer dbClient.DB().Close()

	qparams := c.MustGet(QueryParamsKey).(DeleteImageQueryParams)

	// Deprecating hard deletes for soft deletes
	//stmt, err := dbClient.DB().Prepare("DELETE FROM image WHERE id = '?'")
	stmt, err := dbClient.DB().Prepare("UPDATE image SET deleted = true WHERE id=?")
	if err != nil {
		c.AbortWithError(500, err)
	}
	fmt.Println("qparams[0]", qparams.ID[0])
	idToDelete := qparams.ID[0]
	_, err = stmt.Exec(idToDelete)
	if err != nil {
		c.AbortWithError(500, err)
	}
	/*
		lastID, err := res.LastInsertId()
		if err != nil {
			c.AbortWithError(500, err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			c.AbortWithError(500, err)
		}
		fmt.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
	*/
}
