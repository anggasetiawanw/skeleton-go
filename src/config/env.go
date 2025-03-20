package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config interface {
	GlobalEnvironment() *Envs
	PGSQLDatasource() *PGSQLConnection
}
type config struct {
	envs *Envs
	pgsqlConnect *PGSQLConnection
}

type Envs struct {
	Version         string        `env:"VERSION"`
	Port         	string        `env:"PORT" envDefault:"8080"`
	WriteTimeout    int        	  `env:"WRITE_TIMEOUT"`
	ReadTimeout 	int           `env:"READ_TIMEOUT"`
}

type PGSQLConnection struct {
	PGsqlDB 		string  	`env:"PSQL_DB"`
	PGsqlSchema 	string  	`env:"PSQL_SCHEMA"`
	PGsqlHost 		string  	`env:"PSQL_HOST"`
	PGsqlPort 		int		  	`env:"PSQL_PORT"`
	PGsqlUsername 	string  	`env:"PSQL_USERNAME"`
	PGsqlPassword 	string  	`env:"PSQL_PASSWORD"`
}

func NewConfig() *config{
	return &config{}
}

func (c *config) GlobalEnvironment(envs Envs ) *Envs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	if err := env.Parse(&envs); err != nil {
		panic(err)
	}

	return &envs;
}

func (c *config) PGSQLDatasource(pgsql PGSQLConnection) *PGSQLConnection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	if err := env.Parse(&pgsql); err != nil {
		panic(err)
	}

	return &pgsql;
}