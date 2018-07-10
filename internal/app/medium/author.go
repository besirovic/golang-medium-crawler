package medium

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// SlugResponse represent struct with article slug and author username
type SlugsResponse struct {
	Author string
	Slugs  []string
}

// GetAuthorProfile is responsibe for sending request to author profile
// page and getting data in JSON format
// It receives username as a string
func GetAuthorProfile(username string, c chan SlugsResponse) {
	url := constructMediumAuthorURL(username)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)[16:]

	success := gjson.Get(bodyString, "success").Bool()

	if success != true {
		return
	}

	// user := gjson.GetMany(bodyString, "payload.user.username")
	posts := gjson.Get(bodyString, "payload.references.Post").Map()

	r := SlugsResponse{
		Author: username,
		Slugs:  []string{},
	}

	for _, p := range posts {
		s := gjson.Get(p.String(), "uniqueSlug").String()
		r.Slugs = append(r.Slugs, s)
	}

	c <- r
}
