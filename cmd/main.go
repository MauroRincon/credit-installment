package main

import (
	"credit-installment/internal/handlers"
	"log"
	"net/http"
	"os"
)

// ------- PRDUCTION ---------
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", handlers.FormHandler)
	http.HandleFunc("/calculate", handlers.SubmitHandler)

	// Server
	log.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

// ------- DEVELOPMENT---------
// func main() {
// 	http.HandleFunc("/", handlers.FormHandler)
// 	http.HandleFunc("/calculate", handlers.SubmitHandler)

// 	// Server
// 	log.Println("Run server: http://localhost:3000/")
// 	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
// 		log.Fatalf("Could not start server: %s\n", err.Error())
// 	}
// }
