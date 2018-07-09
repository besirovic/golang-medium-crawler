package main

import (
	"os"

	"github.com/besirovic/medium-crawler/internal/app/medium"
	"github.com/besirovic/medium-crawler/internal/pkg/arango"
	"github.com/besirovic/medium-crawler/internal/pkg/env"
)

func main() {

	// Load ENV variables for development
	env.Load()

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

	slugC := make(chan medium.SlugResponse)
	for _, a := range authors {
		go medium.GetAuthorProfile(a, slugC)
	}

	for s := range slugC {
		go medium.GetArticle(s.Author, s.Slug)
	}
}
