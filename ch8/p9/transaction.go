// Handling transations
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const selOne = "SELECT id, title, content FROM post WHERE ID = $1;"
const insert = "INSERT INTO post(ID, TITLE, CONTENT) VALUES (4, 'Transations Title', 'Transaction Content');"

type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	//
	db := createConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}
	p := Post{}
	// Query in other session/transation
	if err := db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// After commit or rollback the transaction need to recreated.
	tx.Rollback()
}

func createConnection() *sql.DB {
	connStr := "postgres://tn:vnhcmmnmn@10.0.4.200:5432/tn?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
