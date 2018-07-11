package main

import (
	"os"
	"sync"

	"github.com/besirovic/medium-crawler/internal/app/medium"
	"github.com/besirovic/medium-crawler/internal/pkg/arango"
	"github.com/besirovic/medium-crawler/pkg/env"
)

func main() {
	// Load ENV variables for development
	env.Load(".env")

	// Setting up arangoDB connection
	_, _, err := arango.Bootstrap()

	if err != nil {
		os.Exit(1)
		return
	}

	authors := []string{
		"dan_abramov",
		"alexmngn",
		"francescod_ales",
		"nicolascole77",
		"richienorton",
		"JoubranJad",
	}

	slugsResponseC := make(chan medium.SlugsResponse)
	var wg sync.WaitGroup

	for _, a := range authors {
		go medium.GetAuthorProfile(a, slugsResponseC)
	}

	for range authors {
		r := <-slugsResponseC
		wg.Add(len(r.Slugs))
		for _, s := range r.Slugs {
			go medium.GetArticle(r.Author, s, &wg)
		}
	}

	wg.Wait()
}
