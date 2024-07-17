package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
	DB   DB     `yaml:"db"`
}

type DB struct {
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func MustLoad() *Config {
	path := fetchPath()

	_, err := os.Stat(path)
	if err != nil {
		panic("config file not found: " + err.Error())
	}

	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}

	var cfg Config

	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to unmarshal yaml: " + err.Error())
	}

	return &cfg
}

func fetchPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	ifEmpty(path)
	return path
}

func ifEmpty(field string) {
	if field == "" {
		panic(fmt.Sprintf("field %s can not be empty!", field))
	}
}

func (p *Postgres) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", p.User, p.Password, p.Host, p.Port, p.Name)
}
