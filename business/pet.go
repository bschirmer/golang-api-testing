package business

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/bs/a-jumbo-backend-test/config"
	. "github.com/bs/a-jumbo-backend-test/database"
	"github.com/bs/a-jumbo-backend-test/models"
	"github.com/bs/a-jumbo-backend-test/utils"
	"github.com/gorilla/mux"
)

func isValidStatus(status string) bool {
	return status == "available" || status == "pending" || status == "sold"
}

func isValidPet(pet models.Pet) []string {
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

func CreatePetRoute(w http.ResponseWriter, r *http.Request) {
	pet := models.Pet{} //initialize empty pet

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")

	// Parse json request body and use it to set fields on pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// Check pet is valid
	petErr := isValidPet(pet)
	if len(petErr) > 0 {
		w.Write(utils.ProcessResponse(utils.InvalidInputCode, utils.InvalidInput, petErr))
		return
	}

	// Get the next id
	pet.Id = GetMaxId(PetCollection)

	// Save to database
	insertErr := SavePet(pet)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.ServerErrorCode, utils.ServerError, []string{insertErr.Error()}))
		return
	}

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, []string{}))
}

func UpdatePetRoute(w http.ResponseWriter, r *http.Request) {
	pet := models.Pet{} //initialize empty pet

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")

	// Parse json request body and use it to set fields on pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// Check pet is valid
	petErr := isValidPet(pet)
	if len(petErr) > 0 {
		w.Write(utils.ProcessResponse(utils.InvalidInputCode, utils.InvalidInput, petErr))
		return
	}

	// Check is pet exists in db
	existingPet, err := FindPetById(pet.Id)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}
	if existingPet.Id == 0 {
		w.Write(utils.ProcessResponse(utils.NotFoundCode, utils.NotFound, []string{utils.PetNotFound}))
		return
	}

	// Save to database
	updateErr := UpdatePet(pet)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.ServerErrorCode, utils.ServerError, []string{updateErr.Error()}))
		return
	}

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, []string{}))
}

func FindPetByStatusRoute(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query() // Returns a url.Values, which is a map[string][]string
	requestStatuses, _ := vals["status"]

	if len(requestStatuses) == 0 {
		w.Write(utils.ProcessResponse(utils.InvalidCode, utils.InvalidStatus, []string{"No status provided"}))
		return
	}

	var messages []string
	for _, element := range requestStatuses {
		if element != "available" && element != "pending" && element != "sold" {
			messages = append(messages, "Invalid Status: "+element)
		}
	}

	if len(messages) > 0 {
		w.Write(utils.ProcessResponse(utils.InvalidCode, utils.InvalidStatus, messages))
		return
	}

	pets := FindPetsByStatus(requestStatuses)

	responseJson, err := json.Marshal(pets)
	if err != nil {
		panic(err)
	}

	w.Write(responseJson)

}

func FindPetByIdRoute(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r) // Returns a url.Values, which is a map[string][]string
	requestId := vals["petId"]
	if len(requestId) == 0 {
		w.Write(utils.ProcessResponse(utils.InvalidCode, utils.InvalidId, []string{}))
		return
	}

	id, err := strconv.Atoi(requestId)
	pets, err := FindPetById(id)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}

	responseJson, err := json.Marshal(pets)
	if err != nil {
		panic(err)
	}

	w.Write(responseJson)
}

func UpdatePetByIdRoute(w http.ResponseWriter, r *http.Request) {

	pet := models.Pet{} //initialize empty pet

	// get pet id from url
	vals := mux.Vars(r)
	pet.Id, _ = strconv.Atoi(vals["petId"])

	//check it isnt 0
	if pet.Id == 0 {
		w.Write(utils.ProcessResponse(utils.InvalidCode, utils.InvalidInput, []string{}))
		return
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")

	// Parse json request body and use it to set fields on pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// check status is valid
	if !isValidStatus(pet.Status) {
		w.Write(utils.ProcessResponse(utils.InvalidInputCode, utils.InvalidInput, []string{utils.InvalidBody}))
		return
	}

	// update pet
	updateErr := UpdatePet(pet)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.ServerErrorCode, utils.ServerError, []string{updateErr.Error()}))
		return
	}

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, []string{}))

}

func DeletePetByIdRoute(w http.ResponseWriter, r *http.Request) {
	// first thing, check the api key is in the headers
	apiKey := r.Header.Get("api_key")
	if len(apiKey) == 0 {
		//no api key, no deal!
		w.Write(utils.ProcessResponse(utils.ServerErrorCode, utils.ServerError, []string{"Not Authorised"}))
		return
	}

	// check the api key is the same as the one in the config
	var config = Config{}
	config.Read()

	// if you dont have the key, you dont get to delete
	if apiKey != config.ApiKey {
		w.Write(utils.ProcessResponse(utils.ServerErrorCode, utils.ServerError, []string{"Not Authorised"}))
		return
	}

	vals := mux.Vars(r) // Returns a url.Values, which is a map[string][]string
	requestId := vals["petId"]

	log.Print("requestId:" + requestId)

	if len(requestId) == 0 {
		w.Write(utils.ProcessResponse(utils.InvalidCode, utils.InvalidId, []string{}))
		return
	}

	id, _ := strconv.Atoi(requestId)
	if id == 0 {
		w.Write(utils.ProcessResponse(utils.InvalidCode, utils.InvalidId, []string{}))
		return
	}

	// before we do anything stupid, lets check the pet exists
	existingPet, err := FindPetById(id)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}
	if existingPet.Id == 0 {
		w.Write(utils.ProcessResponse(utils.NotFoundCode, utils.NotFound, []string{utils.PetNotFound}))
		return
	}

	err = DeletePetById(id)
	if err != nil {
		w.Write(utils.ProcessResponse(utils.NotFoundCode, utils.PetNotFound, []string{}))
		return
	}

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, []string{}))
}

func UploadImageRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UploadImage: Not Yet Implemented")
}

func DatabaseTesting(w http.ResponseWriter, r *http.Request) {
	err := DeletePetById(1)
	log.Print(err)
}
