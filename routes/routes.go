package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"service-mock/middlewares"
	"service-mock/models"

	"github.com/gorilla/mux"
)

func HandleRequest() {

	configJsonFile, err := os.Open("config.json")

	if err != nil {
		log.Printf("Error occurred when trying to load json file: [%s]", err)
		panic(err)
	}

	configJsonBytes, err := ioutil.ReadAll(configJsonFile)

	if err != nil {
		log.Printf("Error occurred when trying to read json file: [%s]", err)
		panic(err)
	}

	var parsedJson models.Service
	err = json.Unmarshal(configJsonBytes, &parsedJson)

	if err != nil {
		log.Printf("Error occurred when trying parse json file to model: [%s]", err)
		log.Printf("Be sure to keep the json structure the same as the config.json.example")
		panic(err)
	}

	router := mux.NewRouter()
	router.Use(middlewares.ContentTypeMiddleware)

	endpoints := parsedJson.Endpoints
	apiPort := os.Args[1]

	log.Printf("Mocking Service: [%s]", parsedJson.Name)

	for i := 0; i < len(endpoints); i++ {
		endpointName := endpoints[i].Name
		endpointMethod := endpoints[i].Type
		endpointResponse := endpoints[i].Response
		endpointStatusCode := endpoints[i].StatusCode

		router.HandleFunc(endpointName, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(endpointStatusCode)
			json.NewEncoder(w).Encode(endpointResponse)
		}).Methods(endpointMethod)

		log.Printf("Listening to [%s] http://localhost:%s%s...", endpointMethod, apiPort, endpointName)
	}

	log.Fatal(http.ListenAndServe(":"+apiPort, router))
}
