package main

import (
	"com.morrisoncole/katana/ingest"
	"com.morrisoncole/katana/server"
	"log"
)

func main() {
	log.Println("Beginning ingest...")
	ingest.Ingest()
	log.Println("Ingest complete!")

	log.Println("Connecting to Cloud SQL")
	server.Init()

	log.Println("Starting server...")
	server.Serve()
}
