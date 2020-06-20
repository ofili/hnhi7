package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ofili/hnhi7/api"
	_ "html/template"
	"log"
	"net/http"
	"os"
)


func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return p
	}
	return "5000"
}

func main () {
	
	port := getPort()

	// Route definition
	router := mux.NewRouter()
	
	// Render homepage
	router.Handle("/", http.FileServer(http.Dir("./views"))).Methods("GET")

	// 
	router.HandleFunc("/success", api.Pages).Methods("GET") 
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/purchase", api.NotImplemented).Methods("GET")
	router.HandleFunc("/transact", api.Post).Methods("POST")

	fmt.Printf("Listening and serving on port %s.....\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}