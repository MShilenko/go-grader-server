package configs

import (
	"encoding/json"
	"flag"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const serviceName = "go-grader-server"

var options = []option{
	{"config", "string", "", "config file"},

	// {"server.http.host", "string", "127.0.0.1", "server http port"},
	// {"server.http.port", "int", 8086, "server http host"},
	// {"server.http.read_timeout", "int", 40, "server http read timeout"},
	// {"server.http.write_timeout", "int", 40, "server http write timeout"},

	// {"redis.host", "string", "127.0.0.1", "redis host"},
	// {"redis.port", "int", 6379, "redis port"},
	// {"redis.dial_timeout", "int", 10, "redis dial timeout"},
	// {"redis.read_timeout", "int", 30, "redis read timeout"},
	// {"redis.write_timeout", "int", 30, "redis write timeout"},
	// {"redis.pool_size", "int", 10, "redis pool size"},
	// {"redis.pool_timeout", "int", 30, "redis pool timeout"},

	{"postgres.host", "string", "localhost", "postgres host"},
	{"postgres.port", "int", 5432, "postgres port"},
	{"postgres.user", "string", "postgres", "postgres user"},
	{"postgres.password", "string", "postgres", "postgres password"},
	{"postgres.database_name", "string", "grader", "postgres database name"},
	{"postgres.secure", "string", "disable", "postgres SSL support"},
	{"postgres.max_conns_pool", "int", 5, "max number of connections pool postgres"},

	// {"mongo.host", "string", "127.0.0.1", "server http port"},
	// {"mongo.port", "int", 27017, "server http host"},
	// {"mongo.database", "string", "redditclone", "server http host"},
}

// Config - main config struct
type Config struct {
	// Server struct {
	// 	HTTP struct {
	// 		Host         string
	// 		Port         int
	// 		ReadTimeout  int `mapstructure:"read_timeout"`
	// 		WriteTimeout int `mapstructure:"write_timeout"`
	// 	}
	// }
	// Redis struct {
	// 	Host         string
	// 	Port         int
	// 	DialTimeout  int `mapstructure:"dial_timeout"`
	// 	ReadTimeout  int `mapstructure:"read_timeout"`
	// 	WriteTimeout int `mapstructure:"write_timeout"`
	// 	PoolSize     int `mapstructure:"pool_size"`
	// 	PoolTimeout  int `mapstructure:"pool_timeout"`
	// }
	Postgres struct {
		Host         string
		Port         int
		User         string
		Password     string
		DatabaseName string `mapstructure:"database_name"`
		Secure       string
		MaxConnsPool int `mapstructure:"max_conns_pool"`
	}
	// Mongo struct {
	// 	Host     string
	// 	Port     int
	// 	Database string
	// }
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}

// NewConfig new Config instance
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.Read(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Read read parameters for config.
// Read from environment variables, flags or file.
func (c *Config) Read() error {
	for _, o := range options {
		switch o.typing {
		case "string":
			flag.String(o.name, o.value.(string), o.description)
		case "int":
			flag.Int(o.name, o.value.(int), o.description)
		default:
			viper.SetDefault(o.name, o.value)
		}
	}

	viper.SetEnvPrefix(serviceName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic("Read config error: " + err.Error())
	}

	if fileName := viper.GetString("config"); fileName != "" {
		viper.SetConfigFile(fileName)
		viper.SetConfigType("toml")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

// Print print config structure
func (c *Config) Print() error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	log.Println(string(b))
	return nil
}
