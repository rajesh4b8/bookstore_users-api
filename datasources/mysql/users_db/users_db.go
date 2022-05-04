package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
// mysql_username = "mysql_username"
// mysql_password = "mysql_password"
// mysql_host     = "mysql_host"
// mysql_schema   = "mysql_schema"
)

var (
	Client *sql.DB

	// username = os.Getenv(mysql_username)
	// password = os.Getenv(mysql_password)
	// host     = os.Getenv(mysql_host)
	// schema   = os.Getenv(mysql_schema)
)

func init() {
	datasource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		"postgres",
		"dev",
		"127.0.0.1",
		"usersdb",
	)
	var err error
	Client, err = sql.Open("postgres", datasource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connected successfully")
}
