package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUserName     = "mysqlUserName"
	mysqlUserPassword = "mysqlUserPassword"
	mysqlUserHost     = "mysqlUserHost"
	mysqlUserSchema   = "mysqlUserSchema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUserName)
	password = os.Getenv(mysqlUserPassword)
	host     = os.Getenv(mysqlUserHost)
	schema   = os.Getenv(mysqlUserSchema)
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err := sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
