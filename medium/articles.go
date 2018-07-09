package medium

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

type mediumArticle struct {
	auhtor   string
	slug     string
	title    string
	subtitle string
	content  gjson.Result
}

// getArticle is responsible for sending request to article
// page and fetching article data in JSON format
// It receives author username and articleID as strings
func getArticle(username string, slug string) {
	url := constructMediumArticleURL(username, slug)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
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

	postJSON := gjson.GetMany(bodyString, "payload.value.title", "payload.value.content.subtitle", "payload.value.content.bodyModel.paragraphs.#.text")
	a := mediumArticle{
		auhtor:   username,
		slug:     slug,
		title:    postJSON[0].String(),
		subtitle: postJSON[1].String(),
		content:  postJSON[2],
	}

	storeArticle(a)
}

func storeArticle(p mediumArticle) {
	// Store article to ArangoDB
}

// storeArticleLocal is resonsible for storing article to database
// It receives authorID and articleData as strings
func storeArticleLocal(p mediumArticle) {
	fp := filepath.Join(".", "storage", p.auhtor, p.slug)
	fn := fp + ".txt"
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		return
	}

	defer f.Close()

	f.WriteString("Title:" + p.title)
	f.WriteString("\n")
	f.WriteString("Subtitle: " + p.subtitle)
	f.WriteString("\n")
}

// Check if article already exists in database
func checkIfArticleExists(articleID string) {}
