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

var (
	mockDB = map[string]*User{
		"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
	}
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func GetNametest(t *testing.T) {
	log.Println("INITIATING GETByUsernameStreaming TEST")

	echoObj := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	echoContext := echoObj.NewContext(req, rec)

	echoContext.SetPath("/username/:username")
	echoContext.SetParamNames("username")
	echoContext.SetParamValues("andrew alizaga")

	if assert.NoError(t, username.GETByUsernameAndSiteFilteredByFoundStreaming(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	if err != nil {
		t.Logf("ERROR IN GETByUsernameStreaming FUNCTION")
		t.Errorf(err)
	}

	if result == nil {
		t.Logf("ERROR IN GETByUsernameStreaming FUNCTION RESULT")
		t.Errorf("ERROR IN GETByUsernameStreaming FUNCTION RESULT", result)
	}

	t.Log("GET GETByUsernameStreaming FUNCTION WORKS PERFECTLY")
	t.Log(result)

}
