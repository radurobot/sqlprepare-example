// Prepare SQL Queries for Fast API Responses with Gin-Gonic

package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var selectFooById string = `SELECT name FROM foo WHERE id = @id;`

func main() {
	db, err := createAndPopulateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare SQL statement
	stmt, err := db.Prepare(selectFooById)
	if err != nil {
		log.Fatal(err)
	}

	// Create Gin-Gonic router
	router := gin.Default()

	// Handle GET request for /name endpoint
	router.GET("/name", func(c *gin.Context) {
		// Read ID from query string
		id := c.Query("id")

		// Execute SQL statement with given ID and scan the result
		var name string
		// name sql variable id to @id
		err := stmt.QueryRow(sql.Named("id", id)).Scan(&name)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Return name as JSON
		c.JSON(200, gin.H{
			"name": name,
		})
	})

	// Start server on port 8081
	router.Run(":8081")
}

func createAndPopulateDatabase() (*sql.DB, error) {
	// Create in-memory database and populate it with some data
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create table and insert some data
	sqlStmt := `
		CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT);
		INSERT INTO foo(name) VALUES ('foo');
		INSERT INTO foo(name) VALUES ('bar');
	`

	// Execute SQL statements
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}
