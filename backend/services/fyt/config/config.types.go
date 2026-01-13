package config

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
