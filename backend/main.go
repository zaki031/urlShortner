package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/zaki031/shortner/controllers"
	"github.com/zaki031/shortner/database"
)

func main() {
	router := mux.NewRouter()
	database.Connect()
	router.HandleFunc("/", controllers.AddUrl).Methods("POST")
	router.HandleFunc("/{slug}", controllers.GetUrl).Methods("GET")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	corsRouter := corsHandler(router)

	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", corsRouter)
}
