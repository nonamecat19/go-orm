package migrate

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/core/lib/scheme"
	"log"
	"reflect"
)

func PushEntity(config config.ORMConfig, entities []entities.IEntity) {
	//tableConfigs := getAllConfigs(entities)

	switch config.DbDriver {
	case "postgres":
		log.Fatal("postgres push not implemented")
	case "sqlite":
		log.Fatal("sqlite push not implemented")
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
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			continue
		}
		fields = append(fields, scheme.Field{
			Name: dbTag,
			Type: field.Type.Name(),
		})
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
		value := v.Field(i)

		fmt.Printf("%s  %s %v\n", field.Name, field.Type, value)
	}
}
