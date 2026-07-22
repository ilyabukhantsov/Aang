package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	ctx := context.Background()

	urlExample := "postgres://root:mysecretpassword@postgres-db:5432/my_database?sslmode=disable"
	conn, err := pgx.Connect(ctx, urlExample)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("It works")

	defer conn.Close(context.Background())

	const CreateDatabase = `
		CREATE TABLE shorter (
		shorterID INT PRIMARY KEY,
    	original_url TEXT NOT NULL,
    	shorted_url VARCHAR(20) NOT NULL UNIQUE,
		ttl TIMESTAMPTZ NOT NULL
	);
	`

	_, err = conn.Exec(ctx, CreateDatabase)
	if err != nil {
		log.Printf("Failed to create shorter table: %v\n", err)
	} else {
		fmt.Println("Schema initialized: 'shorter' table created successfully.")
	}

	r.Run()
}

func shortURL() {

	return
}
