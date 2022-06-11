package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

func StartDatabase() (*gorm.DB, error) {
	db := "./sherlockgo.db"
	log.Println("Reading database")
	Db, err := gorm.Open("sqlite3", db)
	if err != nil {
		log.Println("[DATABASE][Error]", err)
	}
	// defer Db.Close()
	return Db, err
}
