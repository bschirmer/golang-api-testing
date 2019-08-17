package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bs/a-jumbo-backend-test/business" //How do you make this relative as its internal?
	"github.com/bs/a-jumbo-backend-test/database" //How do you make this relative as its internal?
	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to this api")
}

func handleRequests() {
	// Home
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Headers("Accept", "application/json")

	myRouter.HandleFunc("/", homepage)

	// Pet
	myRouter.HandleFunc("/pet", business.CreatePetRoute).Methods(http.MethodPost)
	myRouter.HandleFunc("/pet", business.UpdatePetRoute).Methods(http.MethodPut) //but how do you pass in params?
	myRouter.HandleFunc("/pet/findByStatus", business.FindPetByStatusRoute).Methods(http.MethodGet)
	myRouter.HandleFunc("/pet/{petId}", business.FindPetByIdRoute).Methods(http.MethodGet)
	myRouter.HandleFunc("/pet/{petId}", business.UpdatePetByIdRoute).Methods(http.MethodPost)
	myRouter.HandleFunc("/pet/{petId}", business.DeletePetByIdRoute).Methods(http.MethodDelete)
	myRouter.HandleFunc("/pet/{petId}/uploadImage", business.UploadImageRoute).Methods(http.MethodPost)

	myRouter.HandleFunc("/test", business.DatabaseTesting).Methods(http.MethodPost)

	myRouter.PathPrefix("/").HandlerFunc(business.NoRoute)

	// Store

	//User

	// Handle any issues
	err := http.ListenAndServe(":8080", myRouter)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func main() {
	database.Init()
	database.CreateTestData()
	handleRequests()
}
