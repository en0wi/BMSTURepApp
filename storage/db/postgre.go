package db

import (
	"database/sql"
)

type DB struct {
	conn *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) CloseDb() error {
	return db.conn.Close()
}
