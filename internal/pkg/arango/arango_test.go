package arango

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("ARANGO_HOST", "http://localhost:8529")
	os.Setenv("ARANGO_USERNAME", "root")
	os.Setenv("ARANGO_PASSWORD", "")
}

func TestConnect(t *testing.T) {
	_, err := connect()

	if err != nil {
		t.Errorf("Error occured while connecting to db %v", err)
	}
}

func TestBootstrap(t *testing.T) {
	_, _, err := Bootstrap()

	if err != nil {
		t.Errorf("Error occured while bootstraping database %v", err)
	}
}

func TestGetClient(t *testing.T) {
	c := GetClient()

	if c == nil {
		t.Errorf("Unable to get client")
	}
}

func TestGetDB(t *testing.T) {
	db := GetDB()

	if db == nil {
		t.Errorf("Unable to get db")
	}
}

func TestGetColl(t *testing.T) {
	coll := GetColl()

	if coll == nil {
		t.Errorf("Unable to get collection")
	}
}

func TestGetSession(t *testing.T) {
	c, db, coll := GetSession()

	if c == nil || db == nil || coll == nil {
		t.Errorf("Unable to retrive session")
	}
}
