package main

import (
	"log"
	"os"

	"github.com/besirovic/medium-crawler/arango"
	"github.com/besirovic/medium-crawler/medium"
	"github.com/besirovic/medium-crawler/utils"
)

func main() {

	// Load ENV variables for development
	utils.LoadENV()

	// Setting up arangoDB connection
	_, _, err := arango.Bootstrap()

	if err != nil {
		log.Panicln("Somethings went wrong", err)
		os.Exit(1)
		return
	}

	medium.Scrape()
}
