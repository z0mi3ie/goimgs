package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// DriverName is the driver name for the sql database
	DriverName = "mysql"
)

// Client is a wrapper around a *sql.DB connection. Don't forget to defer ;)
type Client struct {
	db *sql.DB
}

// NewMySQLClient creates and tests the connection to the
// configured MySQL database and also does a small ping validation.
// This returns an error if one was encountered
func NewMySQLClient(user string, pass string, dbase string) (*Client, error) {
	fmt.Println("NewMySqlClient called")
	fmt.Println("user pass dbase", user, pass, dbase)
	//dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, pass, dbase)
	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", user, pass, dbase)
	fmt.Println("Attempting to connect to database with:", dsn)
	db, err := sql.Open(DriverName, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	c := &Client{
		db: db,
	}
	return c, nil
}

// DB returns the client's database connection
func (c *Client) DB() *sql.DB {
	return c.db
}
