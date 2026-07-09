package models

import (
	"database/sql"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Model struct {
		Id int16
	}
)

func (m *Model) Migrate(model struct{}) error {
	conn, err := m.Connect()

	if err != nil {
		return err
	}
	defer conn.Close()

	query := strings.Builder{}

	t := reflect.TypeOf(model)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		field_parse := strings.ToLower(field.Name)
		type_parse := m.ParseType(field.Type.Name())

		if query.Len() > 0 {
			query.WriteString(", ")
		}
		query.WriteString(field_parse)
		query.WriteString(" ")
		query.WriteString(type_parse)
	}

	return nil
}

func (m *Model) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DatabaseName)
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
