package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB = nil

func StartDatabase(databaseName string) (*gorm.DB, error) {

	var err error = nil

	if DB == nil {

		dbName := "./" + databaseName + ".db"
		log.Println("Reading database")
		DB, err = gorm.Open("sqlite3", dbName)
		if err != nil {
			log.Println("[DATABASE][Error]", err)
		}

	}
	return DB, err
	// defer Db.Close()
}
