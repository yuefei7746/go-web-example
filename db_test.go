package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"testing"
)

func TestPingMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/go_web_example?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}
	err = db.Ping()
	t.Fatal(err)
}
