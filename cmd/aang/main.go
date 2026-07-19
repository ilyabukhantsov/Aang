package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

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
		CREATE TABLE shorters (
		shorterID INT PRIMARY KEY,
    	original_url VARCHAR(150) NOT NULL,
    	shorted_url VARCHAR(100) NOT NULL UNIQUE
	);
	`
	_, err = conn.Exec(ctx, CreateDatabase)
	if err != nil {
		log.Fatalf("Failed to create shorter table: %v\n", err)
	}
	fmt.Println("Schema initialized: 'shorter' table created successfully.")

}
