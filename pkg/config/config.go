package config

// Env returns the value of the environment variable named by the key.
type Env string

// Config is the configuration struct
type Config struct {
	Port        *string `env:"PORT"`
	JWTSecret   string  `env:"JWT_SECRET"`
	DatabaseURL string  `env:"DATABASE_URL"`
	RedisURL    string  `env:"REDIS_URL"`
	Env         string  `env:"ENV"`
	ElasticURL  string  `env:"ELASTIC_URL"`
}

// GetEnv returns the current environment
func (c *Config) GetEnv() Env {
	return Env(c.Env)
}

// Instance is the global configuration
var Instance *Config
