package functional

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pablodz/sherlockgo/internal/endpoints/username"
	"github.com/stretchr/testify/assert"
)

func TestGETByUsernameStreaming(t *testing.T) {
	log.Println("INITIATING GETByUsernameStreaming TEST")

	echoObj := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	echoContext := echoObj.NewContext(req, rec)

	echoContext.SetPath("/username/:username")
	echoContext.SetParamNames("username")
	echoContext.SetParamValues("andrew alizaga")

	if assert.NoError(t, username.GETByUsernameStreaming()(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log("SUCCESS")
	} else {
		t.Errorf("FAILURE")

	}

}

func TestGETByUsernameAndSiteFilteredByFoundStreaming(t *testing.T) {
	log.Println("INITIATING GETByUsernameAndSiteFilteredByFoundStreaming TEST")

	echoObj := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	echoContext := echoObj.NewContext(req, rec)

	echoContext.SetPath("/username/:username/found/:found")
	echoContext.SetParamNames("username", "found")
	echoContext.SetParamValues("andrew alizaga", "true")

	echoContext.SetParamNames("username")
	echoContext.SetParamValues("andrew alizaga")

	if assert.NoError(t, username.GETByUsernameAndSiteFilteredByFoundStreaming()(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log("SUCCESS")
	} else {
		t.Errorf("FAILURE")

	}

}
