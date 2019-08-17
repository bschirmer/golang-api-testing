package database

import (
	"context"
	"log"

	. "a-jumbo-backend-test/config" //How do you make this relative as its internal?

	. "a-jumbo-backend-test/models"

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
	options := options.Find()

	// Sort by `id` field descending
	options.SetSort(bson.D{{"id", -1}})

	// Limit by 1 documents only
	options.SetLimit(1)

	cursor, err := collection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		log.Print(err)
	}

	var id = 0
	// althought this is a foreach, there is only ever 1 document because of the setLimit above
	// the foreach is jsut to access the 1 item in the cursor
	for cursor.Next(context.TODO()) {
		var pet Pet
		// decode the document
		if err := cursor.Decode(&pet); err != nil {
			log.Fatal(err)
		}
		id = pet.Id
	}

	return id + 1
}

func CreateTestData() {

	var tag Tag
	tag.Id = 0
	tag.Name = "Dog Tag"

	var pet Pet
	pet.Category.Id = 0
	pet.Category.Name = "Dog"
	pet.Name = "Doggo"
	pet.PhotoUrls = []string{"1 photo"}
	pet.Tags = []Tag{tag}
	pet.Status = "available"

	// we dont care about these saves
	_, _ = SavePet(pet)

	tag.Id = 0
	tag.Name = "Cat Tag"

	pet.Category.Id = 0
	pet.Category.Name = "Cat"
	pet.Name = "Kitty"
	pet.PhotoUrls = []string{"1 photo"}
	pet.Tags = []Tag{tag}
	pet.Status = "available"

	// we dont care about these saves
	_, _ = SavePet(pet)

	tag.Id = 0
	tag.Name = "Perro Tag"

	pet.Category.Id = 0
	pet.Category.Name = "Perro bonito"
	pet.Name = "Boris"
	pet.PhotoUrls = []string{"1 photo"}
	pet.Tags = []Tag{tag}
	pet.Status = "pending"

	// we dont care about these saves
	_, _ = SavePet(pet)

	tag.Id = 0
	tag.Name = "Gato Tag"

	pet.Category.Id = 0
	pet.Category.Name = "Gato bonito"
	pet.Name = "Gato"
	pet.PhotoUrls = []string{"1 photo"}
	pet.Tags = []Tag{tag}
	pet.Status = "sold"

	// we dont care about these saves
	_, _ = SavePet(pet)
}
