package config

type ORMConfig struct {
	DbDriver string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSLMode  bool
}
