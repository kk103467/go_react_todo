package infrastructure

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// declaration of Database as a GLOBAL VARIABLE
var Database *sql.DB

func ConnectDB() {
	var err error
	Database, err = sql.Open("mysql", "root:password@(localhost:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}

	Database.SetConnMaxLifetime(time.Minute * 3)
	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(10)
}
