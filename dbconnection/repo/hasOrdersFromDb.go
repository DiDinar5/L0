package repo

import (
	"L0/dbconnection"
	"log"
)

func HasOrdersInDB() bool {
	db, err := dbconnection.DbConnection()
	if err != nil {
		log.Panic(err)
		return false
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&count)
	if err != nil {
		log.Panic(err)
		return false
	}

	return count > 0
}
