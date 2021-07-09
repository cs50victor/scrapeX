package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func handleRequests(){
	// := - short variable decleration - doesn't set a type
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func main(){

}