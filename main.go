package main

import (
	_ "github.com/lib/pq"
	"orm/entities"
	"orm/lib/client"
	"orm/lib/config"
	"orm/lib/scheme"
)

func main() {
	userTableSchema := scheme.TableScheme{
		Name: "users",
		Fields: []scheme.Field{
			{Name: "id", Type: "integer"},
			{Name: "name", Type: "varchar(50)"},
			{Name: "email", Type: "varchar(100)"},
			{Name: "password", Type: "varchar(100)"},
		},
	}

	ormConfig := config.ORMConfig{
		DbDriver: "postgres",
		Host:     "localhost",
		Port:     15432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
		Tables: []scheme.TableScheme{
			userTableSchema,
		},
	}

	//ormConfig := config.ORMConfig{
	//	DbDriver: "sqlite",
	//	Host:     "localhost",
	//	Port:     5432,
	//	User:     "postgres",
	//	Password: "root",
	//	DbName:   "orm",
	//	SSLMode:  false,
	//	Tables: []scheme.TableScheme{
	//		userTableSchema,
	//	},
	//}

	db := client.CreateClient(ormConfig)

	user := &entities.User{
		Model: entities.Model{},
		Name:  "Name",
		Email: "Email",
	}

	db.Create(user)

	//err := schema.createTable(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result, err := schema.create(db, "John Doe", "john@example.com", "password123")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
	//
	//users, err := schema.read(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	//
	//result, err = schema.update(db, 1, "Jane Doe", "jane@example.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
	//
	//result, err = schema.delete(db, 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
}
