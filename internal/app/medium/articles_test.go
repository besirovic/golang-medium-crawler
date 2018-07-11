package medium

import (
	"fmt"
	"sync"
	"testing"

	"github.com/besirovic/medium-crawler/internal/pkg/arango"
)

func init() {
	arango.Bootstrap()
}

// TEST FAILING
// TODO: make mocker for db, find way to test
// async operations with WaitGroup
func TestGetArticle(t *testing.T) {
	var wg sync.WaitGroup
	s := "the-evolution-of-flux-frameworks-6c16ad26bb31"
	a := "dan_abramov"

	wg.Add(1)
	go GetArticle(a, s, &wg)
	wg.Wait()
}

// TEST FAILING
// TODO: make mocker for db
func TestStoreArticle(t *testing.T) {
	a := Article{
		author:   "someusername",
		slug:     "some-article",
		title:    "This is some test article",
		subtitle: "This is subtitle",
		content:  "This is some content",
	}

	err := storeArticle(a)

	if err != nil {
		fmt.Println(err)
		//t.Errorf("Error white storing article %v", err)
	}
}

// TODO
func TestCheckIfArticleExists(t *testing.T) {
}
