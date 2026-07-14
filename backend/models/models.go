package models

import (
	"database/sql"
	"fmt"
	"queue-go/utils"
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

	t, _ := utils.ReflectStruct(model)

	table = t.Name()

	columns := []string{"id INTEGER"}

	for i := range t.NumField() {
		field := t.Field(i)

		if !field.Anonymous {
			field_parse := strings.ToLower(field.Name)
			type_parse := m.ParseType(field.Type.Name())

			columns = append(columns, fmt.Sprintf("%s %s", field_parse, type_parse))
		}
	}

	query.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n  %s\n);", strings.ToLower(table), strings.Join(columns, ",\n  ")))

	queryString := query.String()

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

func (m *Model) Columns(model ModelI) ([]string, error) {

	t, _ := utils.ReflectStruct(model)

	table := strings.ToLower(t.Name())

	conn, err := m.Connect()

	defer conn.Close()

	var columns []string

	rows, err := conn.Query(fmt.Sprintf("SELECT name FROM pragma_table_info('%s')", strings.ToLower(table)))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		columns = append(columns, name)
	}

	return columns, nil
}
