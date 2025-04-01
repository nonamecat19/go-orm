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

type ORMConfigYaml struct {
	DB struct {
		DbDriver string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"db"`
		SSLMode  bool   `yaml:"ssl_mode"`
	} `yaml:"db"`
	Migrations struct {
		Path         string `yaml:"path"`
		AddTimestamp bool   `yaml:"add_timestamp"`
	} `yaml:"migrations"`
}
