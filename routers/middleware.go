package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/z0mi3ie/goimgs/config"
	"github.com/z0mi3ie/goimgs/db"
)

// Definitions for gin.Context keys
const (
	MySQLClientKey = "MySQLClientKey"
	QueryParamsKey = "QueryParamsKey"
	ConfigKey      = "ConfigKey"
)

// DeleteImageQueryParams is the query params for a DeleteImage call
type DeleteImageQueryParams struct {
	ID []string `form:"id"`
}

// MySQLClientMiddleware adds a MySQL client to the gin.Context
func MySQLClientMiddleware(c *gin.Context) {
	cfg := c.MustGet(ConfigKey).(*config.Config)
	dbClient, err := db.NewMySQLClient(cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLDatabase)
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

// ConfigMiddleware adds a config struct to the requests context so we can access
// it for the rest of the request
func ConfigMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ConfigKey, cfg)
	}
}
