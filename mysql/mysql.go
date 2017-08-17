package mysql

import (
	"database/sql"
)

const (
	QueryCreateItem     = "create-item"
	QueryDeleteItem     = "delete-item"
	QueryFindItemByID   = "find-item-by-id"
	QueryUpdateItemByID = "update-item-by-id"
)

// Store represents a database connection and a collection of statements
// we will use to interface with our backing MySQL database.
type Store struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

// NewStore will prepare all of our queries on the provided
// database and returns a pointer to a new `mysql.Store` instance.
func NewStore(db *sql.DB) (*Store, error) {
	unprepared := map[string]string{
		QueryCreateItem: `
			INSERT INTO todo.items (title, description)
			VALUES(?, ?);
		`,
		QueryDeleteItem: `
			DELETE FROM todo.items
			WHERE id = ?;
		`,
		QueryFindItemByID: `
			SELECT i.id, i.title, i.description, i.completed, i.created_at, i.updated_at
			FROM todo.items i
			WHERE id = ?;
		`,
		QueryUpdateItemByID: `
			UPDATE todo.items i
			SET 
				i.title = ?,
				i.description = ?,
				i.completed = ?
			WHERE i.id = ?;
		`,
	}

	// prepare all statements to verify syntax
	stmts, err := prepareStmts(db, unprepared)
	if err != nil {
		return nil, err
	}

	s := Store{
		db:    db,
		stmts: stmts,
	}

	return &s, nil
}

// Close closes the database, releasing any open resources.
func (s *Store) Close() error {
	return s.db.Close()
}

// prepareStmts will attempt to prepare each unprepared
// query on the database. If one fails, the function returns
// with an error.
func prepareStmts(db *sql.DB, unprepared map[string]string) (map[string]*sql.Stmt, error) {
	prepared := map[string]*sql.Stmt{}
	for k, v := range unprepared {
		stmt, err := db.Prepare(v)
		if err != nil {
			return nil, err
		}
		prepared[k] = stmt
	}

	return prepared, nil
}
