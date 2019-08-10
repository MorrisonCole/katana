package main

import (
	"com.morrisoncole/katana/ingest"
	"com.morrisoncole/katana/server"
	"log"
)

func main() {
	log.Println("Beginning ingest...")
	jmDict := ingest.Ingest("testdata/dictionary_with_one_entry")
	log.Println("Ingest complete!")

	log.Println("Connecting to Cloud SQL")
	server.Init()

	_, err := server.DB.AddEntry(&server.Entry{Definitions: jmDict.Entries[0].SenseElements[0].Glossary})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")
	server.Serve()
}
