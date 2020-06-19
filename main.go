package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ofili/hnhi7/api"
	_ "html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)




func main () {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(5000)
	}

	// Route definition
	router := mux.NewRouter()
	
	// Render homepage
	router.Handle("/", http.FileServer(http.Dir("./views"))).Methods("GET")

	// 
	router.HandleFunc("/success", api.Pages).Methods("GET") 
	router.PathPrefix("/static/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources/"))))
	router.HandleFunc("/purchase", api.NotImplemented).Methods("GET")
	router.HandleFunc("/transact", api.Post).Methods("POST")

	fmt.Printf("Listening and serving on port %s.....\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}