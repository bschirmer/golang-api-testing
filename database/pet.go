package database

import (
	"context"
	"fmt"
	"log"

	. "github.com/bs/a-jumbo-backend-test/models"
	"go.mongodb.org/mongo-driver/bson"
)

func SavePet(pet Pet) error {
	collection := db.Database(config.Database).Collection(PetCollection)
	_, err := collection.InsertOne(context.TODO(), pet)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdatePet(pet Pet) error {
	collection := db.Database(config.Database).Collection(PetCollection)
	fmt.Print(pet)
	filter := bson.D{{"Id", pet.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"Name", pet.Name},
			{"Status", pet.Status},
			{"PhotoUrls", pet.PhotoUrls},
			{"Category", bson.D{
				{"Id", pet.Category.Id},
				{"Name", pet.Category.Name},
			}},
			{"Tags", pet.Tags},
		}},
	}

	updatedPet, err := collection.UpdateOne(context.TODO(), filter, update)
	fmt.Print(updatedPet)
	return err
}

func FindPetById(id int) (Pet, error) {
	pet, err := dao.FindPetById(id)
	return pet, err
}

func (m *PetStoreDb) FindPetById(id int) (Pet, error) {
	collection := db.Database(config.Database).Collection(PetCollection)

	filter := bson.M{"Id": bson.M{"$eq": id}}
	//filter := bson.M{"Name": bson.M{"$eq": "pet"}}

	var pet Pet
	err := collection.FindOne(context.TODO(), filter).Decode(&pet)

	return pet, err
}

func FindPetsByStatus(statuses []string) []Pet {

	collection := db.Database(config.Database).Collection(PetCollection)

	filter := bson.M{"Status": bson.M{"$in": statuses}}

	// find all documents
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var pets []Pet
	// iterate through all documents
	for cursor.Next(context.TODO()) {
		var pet Pet
		// decode the document
		if err := cursor.Decode(&pet); err != nil {
			log.Fatal(err)
		}
		log.Print(pet)
		pets = append(pets, pet)
	}

	return pets
}
