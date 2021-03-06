package server

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

var DB DictionaryDatabase

func Init() {
	var err error

	DB, err = configureDatastoreDB("katana-249402")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listing Definitions...")
	definitions, err := DB.ListEntries()

	for _, definition := range definitions {
		log.Println(&definition.ID)
	}
	log.Println("Done!")
}

func configureDatastoreDB(projectID string) (DictionaryDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return newDatastoreDB(client)
}
