package handlers

import (
	"cmd/api/main.go/internal/tools"
	"database/sql"

	log "github.com/sirupsen/logrus"
)

func PostToDB() {
	var db *sql.DB
	var err error

	db, err = tools.CreateConnection()
	if err != nil {
		log.Error(err)
		return
	}

	if err := tools.XlsxHandler(db); err != nil {
		log.Error(err)
	}
}
