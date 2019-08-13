package business

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/bs/a-jumbo-backend-test/database"
	"github.com/bs/a-jumbo-backend-test/models"
	"github.com/bs/a-jumbo-backend-test/utils"
	"github.com/gorilla/mux"
)

func checkValidPet(pet models.Pet) []string {
	messages := []string{}
	if pet.Name == "" {
		messages = append(messages, "Pet name is mandatory")
	}
	if len(pet.PhotoUrls) == 0 {
		messages = append(messages, "Photo Urls are mandatory")
	}

	if pet.Status != "available" && pet.Status != "pending" && pet.Status != "sold" {
		messages = append(messages, "Invalid Status: "+pet.Status)
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
	petErr := checkValidPet(pet)
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
	petErr := checkValidPet(pet)
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

	fmt.Print(existingPet)

	// Save to database
	// updateErr := UpdatePet(pet)
	// if err != nil {
	// 	w.Write(utils.ProcessResponse(utils.ServerErrorCode, utils.ServerError, []string{updateErr.Error()}))
	// 	return
	// }

	// return a status ok
	w.Write(utils.ProcessResponse(utils.StatusOKCode, utils.StatusOK, []string{}))
}

func FindPetByStatusRoute(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query() // Returns a url.Values, which is a map[string][]string
	requestStatuses, _ := vals["Status"]

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

	log.Print(requestId)

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
	fmt.Fprint(w, "UpdatePetById: Not Yet Implemented")
}

func DeletePetByIdRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DeletePetById: Not Yet Implemented")
}

func UploadImageRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UploadImage: Not Yet Implemented")
}

func DatabaseTesting(w http.ResponseWriter, r *http.Request) {
	pet, _ := FindPetById(3)
	fmt.Fprint(w, pet)
}
