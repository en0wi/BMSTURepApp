package postgre

import (
	"database/sql"
	"fmt"
)

func New(connectionString string) (*Storage, error) {
	db, err := sql.Open("postgre", connectionString)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
	    id INTEGER PRIMARY KEY
	    name TEXT NOT NULL 
	);
	CREATE INDEX IF NOT EXISTS name ON url(name)
	`)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &Storage{db: db}, nil
}

type Storage struct {
	db *sql.DB
}
