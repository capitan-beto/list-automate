package handlers

import (
	"cmd/api/main.go/internal/tools"
	"database/sql"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func PostToDB(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	var err error

	db, err = tools.CreateConnection()
	if err != nil {
		log.Error(err)
		return
	}
}
