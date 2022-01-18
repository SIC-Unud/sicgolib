package sicgolib

import (
	"database/sql"
	"fmt"
	"log"
)

/*
GetDatabase returns database connected to a programmatically defined database connection.
It throws an error if connection fails in a Fatal manner, because database connection is essential to almost any backend
*/
func GetDatabase(dbAddress string, dbUsername string, dbPassword string, dbName string) *sql.DB {
	log.Printf("INFO GetDatabase database connection: starting database connection process")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUsername, dbPassword, dbAddress, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("ERROR GetDatabase sql open connection fatal error: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ERROR GetDatabase db ping fatal error: %v", err)
	}
	log.Printf("INFO GetDatabase database connection: established successfully with %s\n", dataSourceName)
	return db
}
