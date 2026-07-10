package models

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Model struct {
		Id int16
	}
	ModelI interface {
		GetModel() Model
	}
)

func (m *Model) Migrate(model ModelI) error {
	conn, err := m.Connect()

	if err != nil {
		return err
	}
	defer conn.Close()

	query := strings.Builder{}

	var table string

	t := reflect.TypeOf(model)

	if t.Kind() == reflect.Ptr {
		if t.Elem().Kind() != reflect.Struct {
			return errors.New("It isnt struct")
		}
	} else {
		if t.Kind() != reflect.Struct {
			return errors.New("It isnt struct")
		}
	}

	if t.Kind() == reflect.Ptr {
		table = t.Elem().Name()
	} else {
		table = t.Name()
	}

	columns := []string{"id INTEGER"}

	for i := range t.NumField() {
		field := t.Field(i)

		if !field.Anonymous {
			field_parse := strings.ToLower(field.Name)
			type_parse := m.ParseType(field.Type.Name())

			columns = append(columns, fmt.Sprintf("%s %s", field_parse, type_parse))
		}
	}

	query.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n  %s\n);", table, strings.Join(columns, ",\n  ")))

	queryString := query.String()

	fmt.Println(queryString)
	_, err = conn.Exec(queryString)

	if err != nil {
		return err
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
