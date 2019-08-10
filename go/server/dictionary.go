package server

type Definition struct {
	ID int64
}

type DictionaryDatabase interface {
	ListDefinitions() ([]*Definition, error)

	GetDefinition(id int64) (*Definition, error)

	AddDefinition(b *Definition) (id int64, err error)

	DeleteDefinition(id int64) error

	UpdateDefinition(b *Definition) error

	Close()
}
