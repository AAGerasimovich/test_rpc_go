package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func conn(u, p string) *sql.DB {
	db, err := sql.Open("mysql", u+":"+p+"@/test")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("create table if not exists users 	( Name char(255), UUID  char(32), Date datetime)")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
