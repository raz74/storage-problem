package config

var cfg config

type config struct {
	Postgres postgres `mapstructure:",squash"`
	CsvFile  string   `mapstructure:"CSVFILE"`
	Redis    redis    `mapstructure:",squash"`
}

type postgres struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     int    `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DbName   string `mapstructure:"POSTGRES_DB"`
}

type redis struct {
	MasterName string   `mapstructure:"REDIS_MASTER_NAME"`
	Sentinel   bool     `mapstructure:"REDIS_SENTINEL"`
	Addrs      []string `mapstructure:"REDIS_ADDRS"`
	Password   string   `mapstructure:"REDIS_PASSWORD"`
}

func GetPostgres() *postgres {
	return &cfg.Postgres
}

func GetCsvFile() string {
	return cfg.CsvFile
}
func GetRedis() *redis {
	return &cfg.Redis
}
