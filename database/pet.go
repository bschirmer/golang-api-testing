package database

import (
	"context"
	"log"

	. "golang-api-testing/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SavePet(pet Pet) (primitive.ObjectID, error) {
	collection := db.Database(config.Database).Collection(PetCollection)
	// Get the next id
	pet.Id = GetMaxId(PetCollection)
	insertedResult, err := collection.InsertOne(context.TODO(), pet)
	if err != nil {
		log.Fatal(err)
	}
	insertId, _ := insertedResult.InsertedID.(primitive.ObjectID)
	return insertId, err
}

func UpdatePet(pet Pet) error {
	collection := db.Database(config.Database).Collection(PetCollection)
	filter := bson.M{"id": bson.M{"$eq": pet.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"name", pet.Name},
			{"status", pet.Status},
			{"photoUrls", pet.PhotoUrls},
			{"category", bson.D{
				{"id", pet.Category.Id},
				{"name", pet.Category.Name},
			}},
			{"tags", pet.Tags},
		}},
	}

	// we arent using the updated pet, we only care if the update didnt work
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func FindPetById(id int) (Pet, error) {
	pet, err := dao.FindPetById(id)
	return pet, err
}

func (m *PetStoreDb) FindPetById(id int) (Pet, error) {
	collection := db.Database(config.Database).Collection(PetCollection)

	filter := bson.M{"id": bson.M{"$eq": id}}

	var pet Pet
	err := collection.FindOne(context.TODO(), filter).Decode(&pet)

	return pet, err
}

func FindPetsByStatus(statuses []string) []Pet {

	collection := db.Database(config.Database).Collection(PetCollection)

	filter := bson.M{"status": bson.M{"$in": statuses}}

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
		pets = append(pets, pet)
	}

	return pets
}

func DeletePetById(id int) error {
	collection := db.Database(config.Database).Collection(PetCollection)

	filter := bson.M{"id": bson.M{"$eq": id}}

	// we dont care so much about the result, only if it fails
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}

func FindId(objectID primitive.ObjectID) Pet {
	var pet Pet
	collection := db.Database(config.Database).Collection(PetCollection)

	filter := bson.M{"_id": objectID}
	_ = collection.FindOne(context.TODO(), filter).Decode(&pet)
	return pet
}
