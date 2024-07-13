package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)
import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Field struct {
	Name        string
	Type        string
	Nullability bool
}

type UserScheme struct {
	Fields []Field
}

//func read(db *sql.DB, schema *UserScheme) ([]map[string]interface{}, error) {
//	query := "SELECT * FROM users;"
//	rows, err := db.Query(query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var users []map[string]interface{}
//	for rows.Next() {
//		var user map[string]interface{}
//		err = rows.Scan(&user["id"], &user["name"], &user["email"])
//		if err != nil {
//			return nil, err
//		}
//		users = append(users, user)
//	}
//
//	return users, nil
//}

func update(db *sql.DB, schema *UserScheme, id int, name string, email string) (*sql.Result, error) {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3;"
	result, err := db.Exec(query, name, email, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func delete(db *sql.DB, schema *UserScheme, id int) (*sql.Result, error) {
	query := "DELETE FROM users WHERE id = $1;"
	result, err := db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func create(db *sql.DB, schema *UserScheme, name string, email string, password string) (*sql.Result, error) {
	query := "INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4);"
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(100)
	result, err := db.Exec(query, id, name, email, password)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func createTable(db *sql.DB, schema *UserScheme) error {
	query := "CREATE TABLE IF NOT EXISTS users " + "("
	for _, field := range schema.Fields {
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

func main() {
	connStr := "host=localhost port=5432 user=postgres password=root dbname=orm sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := &UserScheme{
		Fields: []Field{
			{Name: "id", Type: "integer", Nullability: false},
			{Name: "name", Type: "varchar(50)", Nullability: true},
			{Name: "email", Type: "varchar(100)", Nullability: false},
			{Name: "password", Type: "varchar(100)", Nullability: true},
		},
	}

	createTable(db, schema)

	// Create a new user
	result, err := create(db, schema, "John Doe", "john@example.com", "password123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	//// Read all users
	//users, err := read(db, schema)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	//
	// Update an existing user
	result, err = update(db, schema, 1, "Jane Doe", "jane@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Delete a user
	result, err = delete(db, schema, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
