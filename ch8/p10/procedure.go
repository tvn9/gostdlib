// Executing stored procedures and functions
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const call = "SELECT * FROM format_name($1, $2, $3)"
const callMySQL = "CALL simpleproc(?)"

type Result struct {
	Name     string
	Category int
}

func main() {
	db := createConnection()
	defer db.Close()
	r := Result{}

	if err := db.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}

	fmt.Printf("Result is: %+v\n", r)
}

func createConnection() *sql.DB {
	connStr := "postgres://tn:vnhcmmnmn@10.0.4.200:5432/tn?sslmode=disable"
	db, err := sql.Open("postgress", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
