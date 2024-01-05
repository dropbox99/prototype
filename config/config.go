package config

type Config struct {
	Router *Router
}

func NewConfig() *Config {
	injection := NewInjection()

	route := NewRouter(injection)

	return &Config{
		Router: route,
	}
}
