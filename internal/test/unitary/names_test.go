package unitary

import (
	"log"
	"testing"

	username "github.com/pablodz/sherlockgo/internal/endpoints/username"
)

func GetNametest(t *testing.T) {
	log.Println("INITIATING GETNAME TEST")
	name := "andrew alizaga"

	result, err := username.GETByUsername(name)

	if err != nil {
		t.Logf("ERROR IN GETBYUSERNAME FUNCTION")
		t.Errorf(err)
	}

	if result == nil {
		t.Logf("ERROR IN GETBYUSERNAME FUNCTION RESULT")
		t.Errorf("ERROR IN GETBYUSERNAME FUNCTION RESULT", result)
	}

	t.Log("GET GETBYUSERNAME FUNCTION WORKS PERFECTLY")
	t.Log(result)

}
