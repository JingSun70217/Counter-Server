package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var count int

func increaseCounter(w http.ResponseWriter, r *http.Request){
	count = count + 1
	fmt.Fprintf(w, "Counter now: %v", count)
}

func decreaseCounter(w http.ResponseWriter, r *http.Request) {
	count = count - 1
	fmt.Fprintf(w, "Counter now: %v", count)
}

func getCounter(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Counter: %v", count)
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/increaseCounter", increaseCounter)
	router.HandleFunc("/decreaseCounter", decreaseCounter)
	router.HandleFunc("/getCounter", getCounter).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",router))
}

