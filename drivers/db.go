package drivers

import (
	"database/sql"

	"github.com/lib/pq"
)

// DB ...
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	pgSQL, err := pq.ParseURL("elephantsql.com informations")

	db, err := sql.Open("postgres", pgSQL)
	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}
