package models

import (
	"database/sql"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Model struct{}
)

func (m *Model) Migrate(model struct{}) error {
	conn, err := m.Connect()

	if err != nil {
		return err
	}

	query := strings.Builder{}

	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		field_parse := strings.ToLower(field.Name)
		type_parse := m.ParseType(field.Type.Name())

	}
}

func (m *Model) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "app.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *Model) ParseType(typ string) string {

	switch typ {
	case "float32":
		return "float"
	default:
		return "VARCHAR"
	}
}
