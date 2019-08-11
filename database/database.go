package database

import (
	"context"
	"log"

	. "github.com/bs/a-jumbo-backend-test/config" //How do you make this relative as its internal?

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PetStoreDb struct {
	Server   string
	Database string
}

var config = Config{}
var dao = PetStoreDb{}

// var db *mgo.Database
var db *mongo.Client

const (
	PetCollection   = "Pets"
	StoreCollection = "Stores"
	UserCollection  = "Users"
)

// Parse the configuration file 'config.toml', and establish a connection to DB
func Init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Establish a connection to database
func (m *PetStoreDb) Connect() {
	config.Read()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db = client
}

func GetMaxId(collectionName string) int {
	collection := db.Database(config.Database).Collection(PetCollection)
	count, err := collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	return int(count + 1)
}
