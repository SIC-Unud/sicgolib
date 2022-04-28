package sicgolib

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

/*
GetDatabase returns database connected to a programmatically defined database connection.
It throws an error if connection fails in a Fatal manner, because database connection is essential to almost any backend
*/
func GetDatabase(dbAddress string, dbUsername string, dbPassword string, dbName string) *sql.DB {
	log.Printf("INFO GetDatabase database connection: starting database connection process")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUsername, dbPassword, dbAddress, dbName)

	var db *sql.DB
	for {
		db, err := sql.Open("mysql", dataSourceName)
		if pingErr := db.Ping(); err != nil || pingErr != nil {
			if err != nil {
				log.Printf("ERROR GetDatabase sql open connection fatal error: %v\n", err)
			} else if pingErr != nil {
				log.Printf("ERROR GetDatabase db ping fatal error: %v\n", pingErr)
			}
			log.Println("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	log.Printf("INFO GetDatabase database connection: established successfully with %s\n", dataSourceName)
	return db
}
