package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string         `yaml:"env" env-default:"local"`
	StoragePath string         `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration  `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig     `yaml:"grpc"`
	Database    DatabaseConfig `yaml:"database" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type DatabaseConfig struct {
	Host    string `yaml:"host" env-default:"localhost"`
	Port    int    `yaml:"port" env-default:"5432"`
	User    string `yaml:"user" env-default:"postgres"`
	DBName  string `yaml:"dbname" env-default:"sso"`
	SslMode string `yaml:"sslmode" env-default:"disable"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
