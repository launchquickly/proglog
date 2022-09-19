package main

import (
	"log"

	"github.com/launchquickly/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
