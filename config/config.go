package config

import "github.com/spf13/viper"

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

func LoadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("default")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	return nil
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
