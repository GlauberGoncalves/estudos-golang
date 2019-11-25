package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DBDRIVER = "sqlite3"
	DBNAME   = "notas.db"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(DBDRIVER, DBNAME)
	if err != nil {
		panic("failed to connect database")
		return db, err
	}
	return db, err
}
