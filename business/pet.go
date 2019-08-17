package business

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "a-jumbo-backend-test/config"
	. "a-jumbo-backend-test/database"
	. "a-jumbo-backend-test/models"
	"a-jumbo-backend-test/utils"
	"github.com/gorilla/mux"
)

func isValidStatus(status string) bool {
	return status == "available" || status == "pending" || status == "sold"
}

func isValidPet(pet Pet) []string {
	messages := []string{}
	if pet.Name == "" {
		messages = append(messages, "Pet name is mandatory")
	}
	if len(pet.PhotoUrls) == 0 {
		messages = append(messages, "Photo Urls are mandatory")
	}

	// i wanted to use pointer here but not sure how that works in go
	if !isValidStatus(pet.Status) {
		messages = append(messages, "Invalid Status")
	}
	return messages
}

func buildUpdatePetObject(existingPet *Pet, newPet Pet) {
	// check and change things that are updates
	// NB we are saving the existing pet as we know its a proper object
	if newPet.Category != existingPet.Category {
		existingPet.Category = newPet.Category
	}

	if newPet.Name != "" && newPet.Name != existingPet.Name {
		existingPet.Name = newPet.Name
	}
	// In a real solution, I would deal with this, but im not going to waste time on this for this test
	// if newPet.PhotoUrls != existingPet.PhotoUrls {
	// 	existingPet.PhotoUrls = newPet.PhotoUrls
	// }
	// if newPet.Tags != existingPet.Tags {
	// 	existingPet.Tags = newPet.Tags
	// }
	if newPet.Status != "" && newPet.Status != existingPet.Status {
		existingPet.Status = newPet.Status
	}
}

func CreatePetRoute(w http.ResponseWriter, r *http.Request) {
	pet := Pet{} //initialize empty pet

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")

	// Parse json request body and use it to set fields on pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// Check pet is valid
	petErr := isValidPet(pet)
	if len(petErr) > 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidInputCode, utils.InvalidInput, petErr))
		return
	}

	// Save to database
	insertObejectId, insertErr := SavePet(pet)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.ServerErrorCode, utils.ServerError, []string{insertErr.Error()}))
		return
	}

	newPet := FindId(insertObejectId)
	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, newPet))
}

func UpdatePetRoute(w http.ResponseWriter, r *http.Request) {
	pet := Pet{} //initialize empty pet

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")

	// Parse json request body and use it to set fields on pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	//check it isnt 0
	if pet.Id == 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidInput, []string{"Invalid ID supplied"}))
		return
	}

	// Check pet is valid
	petErr := isValidPet(pet)
	if len(petErr) > 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidInputCode, utils.InvalidInput, petErr))
		return
	}

	// Check is pet exists in db
	existingPet, err := FindPetById(pet.Id)
	if err != nil {
		log.Print(err.Error())
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}
	if existingPet.Id == 0 {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.NotFound, []string{utils.PetNotFound}))
		return
	}

	//Check for changed fields
	buildUpdatePetObject(&existingPet, pet)

	// Save to database
	updateErr := UpdatePet(existingPet)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.ServerErrorCode, utils.ServerError, []string{updateErr.Error()}))
		return
	}

	// read the newly updated record for response
	updatedPet, _ := FindPetById(pet.Id)

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, updatedPet))
}

func FindPetByStatusRoute(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query() // Returns a url.Values, which is a map[string][]string
	requestStatuses, _ := vals["status"]

	if len(requestStatuses) == 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidStatus, []string{"No status provided"}))
		return
	}

	var messages []string
	for _, element := range requestStatuses {
		if element != "available" && element != "pending" && element != "sold" {
			messages = append(messages, "Invalid Status: "+element)
		}
	}

	if len(messages) > 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidStatus, messages))
		return
	}

	pets := FindPetsByStatus(requestStatuses)

	if len(pets) == 0 {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.PetNotFound, []string{"No pets found"}))
		return
	}

	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, pets))
}

func FindPetByIdRoute(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r) // Returns a url.Values, which is a map[string][]string
	requestId := vals["petId"]
	if len(requestId) == 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidId, []string{}))
		return
	}

	id, err := strconv.Atoi(requestId)
	pet, err := FindPetById(id)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}

	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, pet))

}

func UpdatePetByIdRoute(w http.ResponseWriter, r *http.Request) {

	pet := Pet{} //initialize empty pet

	// get pet id from url
	vals := mux.Vars(r)
	pet.Id, _ = strconv.Atoi(vals["petId"])

	//check it isnt 0
	if pet.Id == 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidInput, []string{}))
		return
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")

	// Parse json request body and use it to set fields on pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// check status is valid
	if !isValidStatus(pet.Status) {
		w.Write(utils.ProcessBadResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// Check is pet exists in db
	existingPet, err := FindPetById(pet.Id)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}
	if existingPet.Id == 0 {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.NotFound, []string{utils.PetNotFound}))
		return
	}

	//Check for changed fields
	buildUpdatePetObject(&existingPet, pet)

	// update pet
	updateErr := UpdatePet(existingPet)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.ServerErrorCode, utils.ServerError, []string{updateErr.Error()}))
		return
	}

	// read the newly updated record for response
	updatedPet, _ := FindPetById(pet.Id)

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, updatedPet))

}

func DeletePetByIdRoute(w http.ResponseWriter, r *http.Request) {
	// first thing, check the api key is in the headers
	apiKey := r.Header.Get("api_key")
	if len(apiKey) == 0 {
		//no api key, no deal!
		w.Write(utils.ProcessBadResponse(utils.ServerErrorCode, utils.ServerError, []string{"Not Authorised"}))
		return
	}

	// check the api key is the same as the one in the config
	var config = Config{}
	config.Read()

	// if you dont have the key, you dont get to delete
	if apiKey != config.ApiKey {
		w.Write(utils.ProcessBadResponse(utils.ServerErrorCode, utils.ServerError, []string{"Not Authorised"}))
		return
	}

	vals := mux.Vars(r) // Returns a url.Values, which is a map[string][]string
	requestId := vals["petId"]

	if len(requestId) == 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidId, []string{}))
		return
	}

	id, _ := strconv.Atoi(requestId)
	if id == 0 {
		w.Write(utils.ProcessBadResponse(utils.InvalidCode, utils.InvalidId, []string{}))
		return
	}

	// before we do anything stupid, lets check the pet exists
	existingPet, err := FindPetById(id)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}
	if existingPet.Id == 0 {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.NotFound, []string{utils.PetNotFound}))
		return
	}

	err = DeletePetById(id)
	if err != nil {
		w.Write(utils.ProcessBadResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}

	// read the newly updated record for response
	deletedPet, _ := FindPetById(id)

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, deletedPet))
}

func UploadImageRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UploadImage: Not Yet Implemented")
}

func DatabaseTesting(w http.ResponseWriter, r *http.Request) {
	err := GetMaxId(PetCollection)
	log.Print(err)
}
