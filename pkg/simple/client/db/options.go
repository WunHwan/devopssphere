package db

type Options struct {
	DSN string `mapstructure:"dsn,omitempty" description:"database dsn"`
}

func NewDatabaseOptions() *Options {
	return &Options{
		DSN: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
	}
}
