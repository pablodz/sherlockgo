package scraper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/cavaliergopher/grab/v3"
	"github.com/jinzhu/gorm"
	"github.com/pablodz/sherlockgo/internal/models"
)

func LoadData(db *gorm.DB, url string) {
	// download data
	resp, err := grab.Get("./data.json", url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Download saved to", resp.Filename)

	// https://www.golanglearn.com/golang-tutorials/how-to-convert-json-to-csv/
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// Handle error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//initialize struct
	var sites []models.Sites
	// jsonFile's content into 'sites' which we defined above
	err = json.Unmarshal(byteValue, &sites)
	if err != nil {
		log.Println(err)
	}
	// save to database
	for _, site := range sites {
		err := db.Save(&site).Error
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("ALL ADDED")
}
