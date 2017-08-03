package todo

// Store is an interface that describes how the program
// will interact with it's persistent data store.
type Store interface {
	StoreReader
	StoreWriter
	Close() error
}

type StoreReader interface {
	FindItemByID(string) (Item, error)
}

type StoreWriter interface {
	CreateItem(Item) (Item, error)
	DeleteItemByID(string) error
	UpdateItemByID(string, Item) error
}
