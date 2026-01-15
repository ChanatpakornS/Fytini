package config

type Config struct {
	App      AppConfig      `yaml:"app"`
	Services ServicesConfig `yaml:"services"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type ServicesConfig struct {
	Tini TiniServiceConfig `yaml:"tini"`
	Fyt  FytServiceConfig  `yaml:"fyt"`
}

type TiniServiceConfig struct {
	URL string `yaml:"url"`
}

type FytServiceConfig struct {
	URL string `yaml:"url"`
}
