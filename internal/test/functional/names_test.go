package functional

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
	"github.com/pablodz/sherlockgo/internal/utils"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/pablodz/sherlockgo/internal/endpoints/username"
	"github.com/stretchr/testify/assert"
)

func MigrateModels(db *gorm.DB) {
	// Migrate to create tables in database
	db.AutoMigrate(&models.Sites{})
	// db.AutoMigrate(&models.Query{})
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

	echoObj := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	echoContext := echoObj.NewContext(req, rec)

	echoContext.SetPath("/api/v2/username/:username")
	echoContext.SetParamNames("username")
	echoContext.SetParamValues("andrew alizaga")

	if assert.NoError(t, username.GETByUsernameStreaming()(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log("SUCCESS")
	} else {
		t.Errorf("FAILURE")

	}

}
