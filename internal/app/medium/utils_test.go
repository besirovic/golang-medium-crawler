package medium

import (
	"testing"
)

func TestConstructMediumAuthorURL(t *testing.T) {
	u := constructMediumAuthorURL("someusername")
	expU := MediumURL + "@someusername/latest?format=json"

	if u != expU {
		t.Errorf("Invalid author url")
	}
}

func TestConstructMediumArticleURL(t *testing.T) {
	u := constructMediumArticleURL("someuser", "somearticle")
	expU := MediumURL + "@someuser/somearticle?format=json"

	if u != expU {
		t.Errorf("Invalid article url")
	}
}
