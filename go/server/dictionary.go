package server

type Entry struct {
	ID          int64
	Readings    []string
	Definitions []string
}

type DictionaryDatabase interface {
	ListEntries() ([]*Entry, error)

	QueryByReading(word string) ([]Entry, error)

	GetEntry(id int64) (*Entry, error)

	AddEntry(b *Entry) (id int64, err error)

	DeleteEntry(id int64) error

	UpdateEntry(b *Entry) error

	Close()
}
