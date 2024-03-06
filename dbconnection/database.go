package dbconnection

import (
	"L0/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {
	user, host, port, password, dbname := utils.LoadEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
