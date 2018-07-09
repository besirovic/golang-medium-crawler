package arango

import (
	"context"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// Declare arango client, db and collection
var client driver.Client
var db driver.Database
var coll driver.Collection

// Connect function is responsible for connecting app with arango database
func connect() (driver.Client, error) {
	// Get ArangoDB credentials
	arangoHost, arangoUser, arangoPassword := os.Getenv("ARANGO_HOST"), os.Getenv("ARANGO_USERNAME"), os.Getenv("ARANGO_PASSWORD")

	// Connecting to ArangoDB
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{arangoHost},
	})

	if err != nil {
		return nil, err
	}

	// Creating new client
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(arangoUser, arangoPassword),
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

// Bootstrap is responsible for establishing connection, creating db and collections
func Bootstrap() (driver.Database, driver.Collection, error) {
	// Get context and db client
	ctx := context.Background()
	client, err := connect()

	if err != nil {
		return nil, nil, err
	}

	// Get database
	d, err := client.Database(ctx, "medium-crawler")
	db = d

	// Check if database is retrived
	if err != nil {
		return nil, nil, err
	}

	// Check if medium-article collection exists and create one if not
	found, err := db.CollectionExists(ctx, "medium-articles")
	if err != nil {
		return nil, nil, err
	}

	if !found {
		// Create collection
		options := &driver.CreateCollectionOptions{}
		c, err := db.CreateCollection(ctx, "medium-articles", options)

		// Return error if exists
		if err != nil {
			return nil, nil, err
		}
		coll = c
	} else {
		// Retrive collection
		coll, err = db.Collection(ctx, "medium-articles")

		// Return error if exists
		if err != nil {
			return nil, nil, err
		}
	}

	// Return db
	return db, coll, nil
}

// GetClient returns connected client
func GetClient() *driver.Client {
	return &client
}

// GetDB returns database connection
func GetDB() *driver.Database {
	return &db
}

// GetColl returns collection of medium articles
func GetColl() driver.Collection {
	return coll
}

// GetSession returns database and collection
func GetSession() (driver.Client, driver.Database, driver.Collection) {
	return client, db, coll
}
