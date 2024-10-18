package config

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Charset  string
}
