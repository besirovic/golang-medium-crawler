package medium

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/besirovic/medium-crawler/internal/pkg/arango"

	"github.com/tidwall/gjson"
)

// MediumArticle represent struct with needed fields from medium article
type MediumArticle struct {
	auhtor   string
	slug     string
	title    string
	subtitle string
	content  gjson.Result
}

// GetArticle is responsible for sending request to article
// page and fetching article data in JSON format
// It receives author username and articleID as strings
func GetArticle(username string, slug string) {
	url := constructMediumArticleURL(username, slug)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Check if status code is 200
	if resp.StatusCode != http.StatusOK {
		return
	}

	// Getting response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)[16:]

	// Checking if medium success status is OK
	success := gjson.Get(bodyString, "success").Bool()

	if success != true {
		return
	}

	// Getting medium post data
	postJSON := gjson.GetMany(bodyString, "payload.value.title", "payload.value.content.subtitle", "payload.value.content.bodyModel.paragraphs.#.text")
	a := MediumArticle{
		auhtor:   username,
		slug:     slug,
		title:    postJSON[0].String(),
		subtitle: postJSON[1].String(),
		content:  postJSON[2],
	}

	// Storing article to ArangoDB
	storeArticle(a)
}

// storeArticle is responsible for saving article document to ArangoDB
func storeArticle(a MediumArticle) {
	// Get context and collection
	ctx := context.Background()
	coll := arango.GetColl()

	// Convert article struct to map
	p := make(map[string]interface{})
	p["title"] = a.title
	p["subtitle"] = a.subtitle
	p["author"] = a.auhtor
	p["slug"] = a.slug
	p["content"] = a.content.String()

	// Save article as document to db
	coll.CreateDocument(ctx, p)
}

// Check if article already exists in database
func checkIfArticleExists(articleID string) {
	// TODO
}
