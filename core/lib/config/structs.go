package config

type ORMConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSLMode  bool
	Path     string
}
