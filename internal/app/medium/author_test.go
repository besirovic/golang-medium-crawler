package medium

import "testing"

func TestGetAuthorProfile(t *testing.T) {
	u := "dan_abramov"
	c := make(chan SlugsResponse)

	go GetAuthorProfile(u, c)

	r := <-c

	if r.err == true {
		t.Errorf("Error while fetching author profile")
	}
}
