package connections

import (
	"akatsuki/skeleton-go/src/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type PGSQLConnection interface {
	PGSQLDatasource(env *config.PGSQLConnection) (*gorm.DB, error)
	PGSQLClose() error
	PGSQLConnection() *gorm.DB
}

type pgsqlConnect struct {
	db *gorm.DB
}

var myEnv map[string]string
var newLogger logger.Interface

func init() {
	var err error
	myEnv, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error reading environment variables: ", err)
	}

	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
}

func NewPGSQLConnection() (*pgsqlConnect, error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", myEnv["PGSQL_HOST"], myEnv["PGSQL_PORT"], myEnv["PGSQL_USERNAME"], myEnv["PGSQL_DB"], myEnv["PGSQL_PASSWORD"], "disable")

	dts, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   myEnv["PGSQL_SCHEMA"] + ".",
			SingularTable: true,
		},
		Logger:  newLogger,
		NowFunc: time.Now,
	})
	if err != nil {
		return nil, dts.Error
	}
	dbInstance, err := dts.DB()
	if err != nil {
		log.Fatal("Error getting DB instance: ", err)
	}
	dbInstance.SetMaxIdleConns(10)
	dbInstance.SetMaxOpenConns(100)
	dbInstance.SetConnMaxLifetime(time.Hour)
	if err != nil {
		return nil, err
	}

	return &pgsqlConnect{db: dts}, nil
}

func (c *pgsqlConnect) PGSQLConnection() *gorm.DB {
	return c.db
}

func (p *pgsqlConnect) PGSQLClose() error {
	dbInstance, err := p.db.DB()
	if err != nil {
		return err
	}
	_ = dbInstance.Close()
	return nil
}
