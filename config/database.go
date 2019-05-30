package config

type dbconfig struct {
	Username        string
	Password        string
	Host            string
	Port            int
	DBName          string `yaml:"db_name"`
	Driver          string
	Charset         string
	ParseTime       string `yaml:"parse_time"`
	Local           string
	ConnMaxLifeTime int `yaml:"conn_max_life_time"`
	MaxIdleConns    int `yaml:"max_idle_conns"`
	MaxOpenConns    int `yaml:"max_open_conns"`
}
