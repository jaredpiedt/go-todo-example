package mysql

import (
	"errors"

	todo "github.com/jaredpiedt/go-todo-example"
)

// CreateItem will create a new `todo.Item` in the database
func (s *Store) CreateItem(i todo.Item) (todo.Item, error) {
	// Insert the todo item
	result, err := s.stmts[QueryCreateItem].Exec(i.Title, i.Description)
	if err != nil {
		return todo.Item{}, err
	}

	// Get the ID of the new item
	id, err := result.LastInsertId()
	if err != nil {
		return todo.Item{}, err
	}

	// Set the todo Item's id before returning
	i.ID = id

	return i, nil
}

// DeleteItemByID will remove a `todo.Item` with the corresponding id
// from the database.
func (s *Store) DeleteItemByID(id string) error {
	result, err := s.stmts[QueryDeleteItem].Exec(id)
	if err != nil {
		return err
	}

	// Check to make sure an item was deleted
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Make sure we made a delete
	if n == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}

// FindItemByID will return a `todo.Item` with the corresponding id.
func (s *Store) FindItemByID(id string) (todo.Item, error) {
	row := s.stmts[QueryFindItemByID].QueryRow(id)

	var i todo.Item
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err != nil {
		return todo.Item{}, err
	}

	return i, nil
}

// UpdateItemByID will update the specified `todo.Item` in the database.
func (s *Store) UpdateItemByID(id string, i todo.Item) error {
	result, err := s.stmts[QueryUpdateItemByID].Exec(i.Title, i.Description, i.Completed, id)
	if err != nil {
		return err
	}

	// Check to make sure something happened
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Make sure we made an updated
	if n == 0 {
		return errors.New("no rows updated")
	}

	return nil
}
