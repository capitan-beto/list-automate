package main

import (
	"flag"

	"cmd/api/main.go/internal/handlers"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var list = flag.String("list", "none", "list requested for processing")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}

	flag.Parse()
	handlers.Handler(*list)
}
