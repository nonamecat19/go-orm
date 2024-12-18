package migrate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/core/lib/scheme"
	"github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/orm/lib/migrate/postgres"
	"log"
	"reflect"
	"strings"
)

func PushEntity(config config.ORMConfig, entities []entities.IEntity) {
	tableConfigs := getAllConfigs(entities)
	var tablesSql string

	switch config.DbDriver {
	case "postgres":
		tablesSql = postgres.GeneratePostgresTablesSQL(tableConfigs)

	case "sqlite":
		log.Fatal("sqlite push not implemented")

	default:
		log.Fatal(fmt.Printf("dbDriver: %s not supported", config.DbDriver))
	}

	println(tablesSql)
	//executeSql(config, tablesSql)
}

func executeSql(config config.ORMConfig, tablesSql string) {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		config.DbDriver,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
		utils.If(config.SSLMode, "enable", "disable"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	for _, query := range strings.Split(tablesSql, ";") {
		println(query)
		_, err = tx.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func getAllConfigs(entities []entities.IEntity) []scheme.TableScheme {
	var tableConfigs []scheme.TableScheme
	for _, e := range entities {
		tableConfigs = append(tableConfigs, getDbConfig(e))
	}
	return tableConfigs
}

func getDbConfig(data entities.IEntity) scheme.TableScheme {
	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		log.Fatal("data should be a struct")
	}

	var fields []scheme.Field

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Name == "Model" {
			fields = append(fields, scheme.Field{
				//TODO
				Name:        "id",
				Type:        "int64",
				Nullability: false,
			})
		} else {
			dbTag := field.Tag.Get("db")
			if dbTag == "" {
				continue
			}
			typeTag := field.Tag.Get("type")

			if field.Type.Kind() == reflect.Struct {
				println(field.Type.Name())
			}

			fields = append(fields, scheme.Field{
				Name:        dbTag,
				Type:        utils.If(len(typeTag) > 0, typeTag, field.Type.Name()),
				Nullability: field.Tag.Get("nullable") == "true",
			})
		}

	}

	return scheme.TableScheme{
		Name:   t.Name(),
		Fields: fields,
	}
}

func printStruct(data any) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		log.Fatal("Not a struct")
	}

	fmt.Printf("Struct Name: %s\n", t.Name())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s  %s %v\n", field.Name, field.Type, v.Field(i))
	}
}
