package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBConn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=karma-files sslmode=require password=pEOUSii7EkypE4u08bNe host=database-1.ckrzvsguulcs.ap-south-1.rds.amazonaws.com")
	if err != nil {
		log.Fatalln(err)
		return nil, err // Return nil slice and error
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err // Return nil slice and error
	} else {
		log.Println("Successfully Connected")
		return db, nil
	}
}
