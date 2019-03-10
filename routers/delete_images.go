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

	stmt, err := dbClient.DB().Prepare("DELETE FROM image WHERE id='?'")
	if err != nil {
		c.AbortWithError(500, err)
	}
	res, err := stmt.Exec(qparams.ID)
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
}
