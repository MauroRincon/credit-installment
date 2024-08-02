package main

import (
	"credit-installment/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.FormHandler)
	http.HandleFunc("/calculate", handlers.SubmitHandler)
	//creditCalculation()

	// Server
	log.Println("Run server: http://localhost:3000/")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
