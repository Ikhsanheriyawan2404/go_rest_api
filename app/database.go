package app

import (
	"database/sql"
	"golang_rest_api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3360)/golang_rest")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
