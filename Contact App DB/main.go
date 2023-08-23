package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"login/controller"
)

func main() {
	// Recover from any panics and log the error
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()

	HandleFunction()
}

func HandleFunction() {
	// Define allowed headers, origins, and methods for CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                     // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}, // Allowing only get, just an example
	})

	// Create a new Gorilla mux router
	router := mux.NewRouter().StrictSlash(true)

	// Define the base path for the API
	router = router.PathPrefix("/api/v1/contactApp").Subrouter()

	// Define routes for the login functionality
	router.HandleFunc("/login", controller.Login).Methods("POST")

	// Routes for user management
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", controller.CreateUser).Methods("POST")
	userRouter.HandleFunc("/get/{id}", controller.GetUserById).Methods("GET")
	userRouter.HandleFunc("/getall/{page}", controller.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/update/{id}", controller.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/delete/{id}", controller.DeleteUser).Methods("DELETE")

	// Routes for contact management
	contactRouter := router.PathPrefix("/{userid}/contacts").Subrouter()
	contactRouter.HandleFunc("/", controller.CreateContact).Methods("POST")
	contactRouter.HandleFunc("/get/{id}", controller.GetContactById).Methods("GET")
	contactRouter.HandleFunc("/getall/{page}", controller.FindAllContacts).Methods("GET")
	contactRouter.HandleFunc("/update", controller.UpdateContact).Methods("PUT")
	contactRouter.HandleFunc("/delete", controller.DeleteContact).Methods("DELETE")

	// Routes for contact info management
	contactInfoRouter := router.PathPrefix("/{userid}/contactinfo").Subrouter()
	contactInfoRouter.HandleFunc("/", controller.CreateContactInfo).Methods("POST")
	contactInfoRouter.HandleFunc("/get", controller.GetContactInfoById).Methods("GET")
	contactInfoRouter.HandleFunc("/getall/{page}", controller.FindAllContactInfo).Methods("GET")
	contactInfoRouter.HandleFunc("/update", controller.UpdateContactInfo).Methods("PUT")
	contactInfoRouter.HandleFunc("/delete", controller.DeleteContactInfo).Methods("DELETE")

	// Start the server on localhost:3000
	log.Printf("Server Live on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", c.Handler(router)))
}
