package tools

import (
	"database/sql"
	"fmt"
	"os"

	"cmd/api/main.go/models"

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

func AddProduct(p *models.Product, db *sql.DB) error {
	id := p.ID
	desc := p.Desc
	price := p.Price
	date := p.Date

	query := "REPLACE INTO products_db (id, item_desc, price, date) VALUES(?, ?, ?, ?)"
	_, err := db.Exec(query, &id, &desc, &price, &date)
	if err != nil {
		return err
	}

	fmt.Println("ID: " + id + " added!")
	return nil
}
