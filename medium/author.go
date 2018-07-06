package medium

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

type slugResponse struct {
	author string
	slug   string
}

// getAuthorProfile is responsibe for sending request to author profile
// page and getting data in JSON format
// It receives username as a string
func getAuthorProfile(username string, c chan slugResponse) {
	createAuthorDir(username)

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
		c <- slugResponse{
			author: username,
			slug:   s,
		}
	}
}

func createAuthorDir(u string) {
	// Check if author's directory exists
	d := filepath.Join(".", "storage", u)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		os.MkdirAll(d, os.ModePerm)
	}
}