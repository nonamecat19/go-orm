mssqlConfig := config.ORMConfig{
	Host:     db.Host,
	Port:     db.Port,
	User:     db.User,
	Password: db.Password,
	DbName:   db.DbName,
	SSLMode:  db.SSLMode,
}

mssqlAdapter := adaptermssql.AdapterMSSQL{}

dbClient := client2.CreateClient(mssqlConfig, mssqlAdapter)

var users []entities.User

err := querybuilder.CreateQueryBuilder(dbClient).
	Where("name = ?", name).
	AndWhere("age >= 18").
	Preload("orders").
	Preload("role").
	OrderBy("id DESC").
	Limit(pageSize).
	Offset(pageSize * (page - 1)).
	Debug().
	FindMany(&users)
