package scheme

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type Field struct {
	Name        string
	Type        string
	Nullability bool
}

type TableScheme struct {
	Name   string
	Fields []Field
}

func (ts TableScheme) Read(db *sql.DB) ([]map[string]interface{}, error) {
	query := "SELECT * FROM users;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var user map[string]interface{}
		err = rows.Scan(user["id"], user["name"], user["email"])
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ts TableScheme) Update(db *sql.DB, id int, name string, email string) (*sql.Result, error) {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3;"
	result, err := db.Exec(query, name, email, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ts TableScheme) Delete(db *sql.DB, id int) (*sql.Result, error) {
	query := "DELETE FROM users WHERE id = $1;"
	result, err := db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ts TableScheme) Create(db *sql.DB, name string, email string, password string) (*sql.Result, error) {
	query := "INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4);"
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	result, err := db.Exec(query, id, name, email, password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ts TableScheme) CreateTable(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS " + ts.Name + " ("
	for _, field := range ts.Fields {
		query += fmt.Sprintf("%s %s%s, ", field.Name, field.Type,
			ifNotNullable(field.Nullability))
	}
	query = query[:len(query)-2] + ");"

	_, err := db.Exec(query)
	return err
}

func ifNotNullable(nullable bool) string {
	if !nullable {
		return " NOT NULL"
	}
	return ""
}
