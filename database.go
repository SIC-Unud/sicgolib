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

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("ERROR GetDatabase sql open connection fatal error: %v", err)
		for {
			log.Println("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			db, err = sql.Open("mysql", dataSourceName)
			if err == nil {
				break
			}
		}
	}
	if err = db.Ping(); err != nil {
		log.Printf("ERROR GetDatabase db ping fatal error: %v", err)
		for {
			log.Println("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			db, err = sql.Open("mysql", dataSourceName)
			err2 := db.Ping()
			if err == nil && err2 == nil {
				break
			}
		}
	}
	log.Printf("INFO GetDatabase database connection: established successfully with %s\n", dataSourceName)
	return db
}
