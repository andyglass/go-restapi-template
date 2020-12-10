package db

import (
	"database/sql"

	"github.com/andyglass/go-restapi-template/logger"
)

var (
	log = logger.Logger
	// DB global
	DB *sql.DB
)

// Get DB conn
func Get(connStr string) *sql.DB {
	// sql.Register("postgreshook", sqlhooks.Wrap(&pq.Driver{}, loghooks.New()))

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var db *sql.DB

	// if err = db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }

	// dbSetMaxIdleConns := 5
	// dbSetMaxOpenConns := 10

	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// db.SetMaxIdleConns(dbSetMaxIdleConns)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// db.SetMaxOpenConns(dbSetMaxOpenConns)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// db.SetConnMaxLifetime(time.Hour)

	return db
}
