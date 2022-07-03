package functional

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/endpoints/username"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
	"github.com/pablodz/sherlockgo/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	// Migrate to create tables in database
	db.AutoMigrate(&models.Sites{})
	db.AutoMigrate(&models.UsernameResponse{})
}

func TestGETByUsernameStreaming(t *testing.T) {
	log.Println("INITIATING GETByUsernameStreaming TEST")
	db, err := database.GetDB()
	if err != nil {
		t.Errorf("DATABASE ERROR")
	}
	// migrate
	MigrateModels(db)
	// download json from sherlock
	scraper.LoadData(db, utils.TestUrl)

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v2/username/:username")
	c.SetParamNames("username")
	c.SetParamValues("andrew")

	// Assertions
	if assert.NoError(t, username.GETByUsernameStreaming()(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	} else {
		t.Errorf("ERROR")
	}
}
