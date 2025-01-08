package client

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/core/lib/scheme"
	"log"
	"reflect"
	"strings"
)

type Tables = map[string]scheme.TableScheme

type DbClient struct {
	db     *sql.DB
	config config.ORMConfig
	tables Tables
}

func CreateClient(config config.ORMConfig) DbClient {
	var connStr string
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	db, err := sql.Open(config.DbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//}(db)

	tableMap := make(Tables)

	return DbClient{
		db:     db,
		config: config,
		tables: tableMap,
	}
}

func (dc DbClient) GetDb() *sql.DB {
	return dc.db
}

func (dc DbClient) GetConfig() config.ORMConfig {
	return dc.config
}

func (dc DbClient) GetTables() Tables {
	return dc.tables
}

//func (dc DbClient) Read() ([]entities2.User, error) {
//	query := "SELECT name FROM users;"
//	rows, err := dc.db.Query(query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var users []entities2.User
//	for rows.Next() {
//		var user entities2.User
//		err = rows.Scan(&user.Name)
//		if err != nil {
//			return nil, err
//		}
//		users = append(users, user)
//	}
//
//	return users, nil
//}

func (dc DbClient) Read(entity interface{}) ([]interface{}, error) {
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("entity must be a pointer to a struct")
	}

	entityType = entityType.Elem()
	tableName := ""
	if tableNameMethod, ok := reflect.New(entityType).Interface().(interface{ TableName() string }); ok {
		tableName = tableNameMethod.TableName()
	} else {
		return nil, errors.New("entity struct must implement TableName() string method")
	}

	var fieldNames []string
	for i := 0; i < entityType.NumField(); i++ {
		fieldTag := entityType.Field(i).Tag.Get("db")
		if fieldTag != "" {
			fieldNames = append(fieldNames, fieldTag)
		}
	}

	query := "SELECT " + joinFields(fieldNames) + " FROM " + tableName
	println(query)
	rows, err := dc.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		entityValue := reflect.New(entityType).Elem()
		var fields []interface{}
		// 1 to avoid id in entity
		for i := 1; i < entityType.NumField(); i++ {
			fields = append(fields, entityValue.Field(i).Addr().Interface())
		}

		if err := rows.Scan(fields...); err != nil {
			return nil, err
		}

		results = append(results, entityValue.Interface())
	}

	return results, nil
}

func joinFields(fields []string) string {
	return "" + strings.Join(fields, ", ") + ""
}

func (dc DbClient) Update(id int, name string, email string) (*sql.Result, error) {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3;"
	result, err := dc.db.Exec(query, name, email, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (dc DbClient) Delete(id int) (*sql.Result, error) {
	query := "DELETE FROM users WHERE id = $1;"
	result, err := dc.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (dc DbClient) Create(entity entities.IEntity) {
	fmt.Print(entity)
	v := reflect.ValueOf(entity).Elem()
	fmt.Print(v)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("Field name: %s, value: %v\n, name: %s ", v.Type().Field(i).Name, field.Interface(), v.Type().Field(i).Tag.Get("db"))
	}

	//tableName := entity.TableName()
	//query := "INSERT INTO " + tableName + " (id, name, email, password) VALUES ($1, $2, $3, $4);"
	//rand.Seed(time.Now().UnixNano())
	//id := rand.Intn(100)
	//result, err := dc.db.Exec(query, id, name, email, password)
	//if err != nil {
	//	return nil, err
	//}
	//return &result, nil
}

//func (ts TableScheme) CreateTable(db *sql.DB) error {
//	query := "CREATE TABLE IF NOT EXISTS " + ts.Name + " ("
//	for _, field := range ts.Fields {
//		query += fmt.Sprintf("%s %s%s, ", field.Name, field.Type,
//			ifNotNullable(field.Nullability))
//	}
//	query = query[:len(query)-2] + ");"
//
//	_, err := db.Exec(query)
//	return err
//}

func ifNotNullable(nullable bool) string {
	if !nullable {
		return " NOT NULL"
	}
	return ""
}
