package server

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

type datastoreDB struct {
	client *datastore.Client
}

var _ DictionaryDatabase = &datastoreDB{}

func newDatastoreDB(client *datastore.Client) (DictionaryDatabase, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &datastoreDB{
		client: client,
	}, nil
}

func (db *datastoreDB) datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey("Definition", id, nil)
}

func (db *datastoreDB) ListDefinitions() ([]*Definition, error) {
	ctx := context.Background()
	definitions := make([]*Definition, 0)
	q := datastore.NewQuery("Definition").
		Order("Title")

	keys, err := db.client.GetAll(ctx, q, &definitions)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list definitions: %v", err)
	}

	for i, k := range keys {
		definitions[i].ID = k.ID
	}

	return definitions, nil
}

func (db *datastoreDB) GetDefinition(id int64) (*Definition, error) {
	ctx := context.Background()
	k := db.datastoreKey(id)
	definition := &Definition{}
	if err := db.client.Get(ctx, k, definition); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Definition: %v", err)
	}
	definition.ID = id
	return definition, nil
}

func (db *datastoreDB) AddDefinition(b *Definition) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("Definition", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Definition: %v", err)
	}
	return k.ID, nil
}

func (db *datastoreDB) DeleteDefinition(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Definition: %v", err)
	}
	return nil
}

func (db *datastoreDB) UpdateDefinition(b *Definition) error {
	ctx := context.Background()
	k := db.datastoreKey(b.ID)
	if _, err := db.client.Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update Definition: %v", err)
	}
	return nil
}

func (db *datastoreDB) Close() {
}
