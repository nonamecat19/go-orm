package config

type StudioConfig struct {
	ServerAddr string
}

func NewConfig() StudioConfig {
	return StudioConfig{
		ServerAddr: ":8080",
	}
}
