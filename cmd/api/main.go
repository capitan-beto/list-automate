package main

import (
	"flag"

	"cmd/api/main.go/internal/handlers"
)

var list = flag.String("list", "none", "list requested for processing")

func main() {
	flag.Parse()
	handlers.Handler(*list)
}
