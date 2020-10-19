package database

import (
	"log"
	"os"
	"osvaldoabel/users-api/src/domain"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func init() {
	godotenv.Load("./../../.env")
}

func NewDb() *Database {
	return &Database{}
}

/**
	Description: Create new Connection for test env.
	Return: connection *gorm.DB
**/
func NewDbTest() *gorm.DB {
	db := NewDb()
	db.Env = "test"
	db.DbTypeTest = "sqlite3"
	db.DsnTest = ":memory:"
	db.AutoMigrateDb = true
	db.Debug = true

	conn, err := db.Connect()

	if err != nil {
		log.Fatalf("[%v] db error: %v", db.Env, err)
	}

	return conn
}

/**
	Description: Connect to the database, acording
		to parameters set in .env file

	Return: connection *gorm.DB
**/
func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	dbType := d.DbType
	dsn := d.Dsn

	if d.Env == "test" {
		dbType = d.DbTypeTest
		dsn = d.DsnTest
	}

	d.Db, err = gorm.Open(dbType, dsn)

	if err != nil {
		return nil, err
	}

	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		return nil, err
	}

	d.Db.LogMode(debugMode)

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.User{})
	}

	return d.Db, nil
}
