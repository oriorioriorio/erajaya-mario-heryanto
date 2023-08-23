package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(mysql:3306)/erajaya?charset=utf8mb4&loc=Local&parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
