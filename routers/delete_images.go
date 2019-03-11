package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/db"
)

// DeleteImages removes desired images through the query parameter "id"
func DeleteImages(c *gin.Context) {
	dbClient := c.MustGet(MySQLClientKey).(*db.Client)
	defer dbClient.DB().Close()

	qparams := c.MustGet(QueryParamsKey).(DeleteImageQueryParams)

	for _, qpID := range qparams.ID {
		// Deprecating hard deletes for soft deletes
		//stmt, err := dbClient.DB().Prepare("DELETE FROM image WHERE id = '?'")
		stmt, err := dbClient.DB().Prepare("UPDATE image SET deleted = true WHERE id=?")
		if err != nil {
			c.AbortWithError(500, err)
		}
		_, err = stmt.Exec(qpID)
		if err != nil {
			c.AbortWithError(500, err)
		}
	}
}
