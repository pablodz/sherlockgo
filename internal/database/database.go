package database

import (
	"log"

	"github.com/pablodz/sherlockgo/internal/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func GetDB() (*gorm.DB, error) {

	var err error = nil

	if DB == nil {

		dbName := "./" + utils.MainDbName + ".db"
		log.Println("Reading database")
		DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Println("[DATABASE][Error]", err)
		}

	}
	return DB, err
	// defer Db.Close()
}
