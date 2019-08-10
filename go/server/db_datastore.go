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
	return datastore.IDKey("Entry", id, nil)
}

func (db *datastoreDB) ListEntries() ([]*Entry, error) {
	ctx := context.Background()
	entries := make([]*Entry, 0)
	q := datastore.NewQuery("Entry").
		Order("Title")

	keys, err := db.client.GetAll(ctx, q, &entries)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list entries: %v", err)
	}

	for i, k := range keys {
		entries[i].ID = k.ID
	}

	return entries, nil
}

func (db *datastoreDB) GetEntry(id int64) (*Entry, error) {
	ctx := context.Background()
	k := db.datastoreKey(id)
	entry := &Entry{}
	if err := db.client.Get(ctx, k, entry); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Entry: %v", err)
	}
	entry.ID = id
	return entry, nil
}

func (db *datastoreDB) AddEntry(b *Entry) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("Entry", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Entry: %v", err)
	}
	return k.ID, nil
}

func (db *datastoreDB) DeleteEntry(id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(id)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Entry: %v", err)
	}
	return nil
}

func (db *datastoreDB) UpdateEntry(b *Entry) error {
	ctx := context.Background()
	k := db.datastoreKey(b.ID)
	if _, err := db.client.Put(ctx, k, b); err != nil {
		return fmt.Errorf("datastoredb: could not update Entry: %v", err)
	}
	return nil
}

func (db *datastoreDB) Close() {
}
