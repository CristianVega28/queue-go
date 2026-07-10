package models

import "database/sql"

const DatabaseName = "queue-go.db"

func Setup() error {
	db, err := sql.Open("sqlite3", DatabaseName)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Ping()
}
