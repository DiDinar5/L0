package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, int, string, string) {
	err := godotenv.Load("C:/Users/DinarKhayrutdinov/source/repos/L0/.env")
	if err != nil {
		HandleError(err, true)
	}

	port, _ := strconv.Atoi(os.Getenv("port"))
	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	return user, host, port, password, dbname
}
