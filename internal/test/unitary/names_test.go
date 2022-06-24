package unitary

import (
	"log"
	"testing"

	"github.com/labstack/echo"
	username "github.com/pablodz/sherlockgo/internal/endpoints/username"
)

func GetNametest(t *testing.T) {
	log.Println("INITIATING GETByUsernameStreaming TEST")
	name := "andrew alizaga"

	echo := echo.New()
	echoC := echo.NewContext()

	echoContext.AcquireContext().Param("username")
	echoContext.AcquireContext().ParamValues()

	result, err := username.GETByUsernameStreaming(name)

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
