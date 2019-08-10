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

	firstReading := jmDict.Entries[0].ReadingElements[0].KReb
	definitions := jmDict.Entries[0].SenseElements[0].Glossary
	_, err := server.DB.AddEntry(&server.Entry{Definitions: definitions, Readings: []string{firstReading}})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")
	server.Serve()
}
