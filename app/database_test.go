package app

import (
	"database/sql"
	"golang_rest_api/helper"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func ConnDb() {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/golang_rest")

	helper.PanicIfError(err)

	insert, err := db.Query("INSERT INTO categories(name) VALUE('ikhsan')")
	// if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func TestDB(t *testing.T) {
	ConnDb()
}
