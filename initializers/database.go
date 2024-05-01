package initializers

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func LoadDatabase() {
	dbURI := os.Getenv("MONGODB_URI")

	if dbURI == "" {
		panic("MONGODB_URI is not set")
	}

	// Connect to the database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
