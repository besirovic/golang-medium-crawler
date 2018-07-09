package medium

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// SlugResponse represent struct with article slug and author username
type SlugResponse struct {
	Author string
	Slug   string
}

// GetAuthorProfile is responsibe for sending request to author profile
// page and getting data in JSON format
// It receives username as a string
func GetAuthorProfile(username string, c chan SlugResponse) {
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

	for _, p := range posts {
		s := gjson.Get(p.String(), "uniqueSlug").String()
		c <- SlugResponse{
			Author: username,
			Slug:   s,
		}
	}
}
