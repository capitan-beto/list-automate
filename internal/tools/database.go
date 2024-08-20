package tools

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func CreateConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Error(err)
	}

	db.SetMaxOpenConns(5)
	fmt.Println("Connected!")
	return db, err
}
