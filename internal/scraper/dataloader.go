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
	var sites map[string]models.Sites
	// jsonFile's content into 'sites' which we defined above
	// fmt.Println(string(byteValue))
	err = json.Unmarshal(byteValue, &sites)
	if err != nil {
		log.Println(err)
	}
	// save to database
	for siteName, siteProps := range sites {
		// Check if already exists
		var siteTmp []models.Sites
		db.Find(&siteTmp, &models.Sites{Sitename: siteName})
		if len(siteTmp) != 0 {
			log.Println("ALREADY ADDED: ", siteName)
			continue
		} else {
			log.Println("ADDED NOW: ", siteName)
		}

		db.Create(&models.Sites{
			Sitename:          siteName,
			ErrorType:         siteProps.ErrorType,
			ErrorMessage:      siteProps.ErrorMessage,
			URLDomain:         siteProps.URLDomain,
			URLFormat:         siteProps.URLFormat,
			UsernameClaimed:   siteProps.UsernameClaimed,
			UsernameUnclaimed: siteProps.UsernameUnclaimed,
		})
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("ALL ADDED")
}
