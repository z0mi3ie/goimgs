package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/config"
	"github.com/z0mi3ie/goimgs/db"
)

// Definitions for gin.Context keys
const (
	MySQLClientKey = "MySQLClientKey"
	QueryParamsKey = "QueryParamsKey"
)

// DeleteImageQueryParams is the query params for a DeleteImage call
type DeleteImageQueryParams struct {
	ID []string `form:"id"`
}

// MySQLClientMiddleware adds a MySQL client to the gin.Context
func MySQLClientMiddleware(c *gin.Context) {
	fmt.Println("Got it!")
	dbClient, err := db.NewMySQLClient(config.MySQLUser, config.MySQLPassword, config.MySQLDatabase)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.Set(MySQLClientKey, dbClient)
}

// DeleteImageQueryParamsMiddleware validates and extracts a DeleteImageQueryParams request
func DeleteImageQueryParamsMiddleware(c *gin.Context) {
	var params DeleteImageQueryParams
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.Set(QueryParamsKey, params)
}

// CORSHeaderMiddleware handles setting up the CORS header for responses
func CORSHeaderMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
}
