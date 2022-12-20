package main

import (
	"log"
	"service-mock/routes"
)

func main() {
	log.Printf("Initiating service-mock API...")
	routes.HandleRequest()
}
