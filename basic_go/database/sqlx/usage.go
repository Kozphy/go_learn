package sqlx

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CRETE TABLE person (
	first_name text,
	last_name text,
	email text
);

CREATE TABLE place (
	country text,
	city text NULL,
	telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func Exec_Connect() {
	var db *sqlx.DB
	var err error
	db_name := "mysql"
	db_url := "root:root@tcp(127.0.0.1:3306)"
	database := "/hello"

	db, err = sqlx.Connect(db_name, db_url+database)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
}
