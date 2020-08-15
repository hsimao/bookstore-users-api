package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUsername = "mysqlUsersUsername"
	mysqlUsersPassword = "mysqlUsersPassword"
	mysqlUsersHost     = "mysqlUsersHost"
	mysqlUsersSchema   = "mysqlUsersSchema"
)

var (
	Client *sql.DB
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv(mysqlUsersUsername)
	password := os.Getenv(mysqlUsersPassword)
	host := os.Getenv(mysqlUsersHost)
	schema := os.Getenv(mysqlUsersSchema)

	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	fmt.Println("datasourceName ", datasourceName)

	var err error

	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
