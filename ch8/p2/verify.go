// Validating the connection
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://admin:password@108.0.67.80:5432/?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping OK.")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	err = db.PingContext(ctx)
	defer cancel()
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	// Verify the connection
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.PingContext(context.Background())
	err = conn.PingContext(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Ping OK.")
}
