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

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to use counter serve. \n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "- To increase counter, please navigate to /increaseCounter \n")
	fmt.Fprintf(w, "- To decrease counter, please navigate to /decreaseCounter \n")
	fmt.Fprintf(w, "- To show the current counter, please navigate to /getCounter \n")
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/increaseCounter", increaseCounter)
	router.HandleFunc("/decreaseCounter", decreaseCounter)
	router.HandleFunc("/getCounter", getCounter).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",router))
}

