package todo

// Store is an interface that describes how the program
// will interact with it's persistent data store.
type Store interface {
	CreateItem(Item) (Item, error)
	DeleteItemByID(string) error
	FindItemByID(string) (Item, error)
	UpdateItemByID(string, Item) error

	Close() error
}
