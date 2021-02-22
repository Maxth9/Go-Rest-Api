package main

import (
	"fmt"
	"log"
	"net/http"
)

func landingPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Endpoint Hit!")
}

func handleRequest(){
	http.HandleFunc("/", landingPage)
	log.Fatalf(http.ListenAndServe(":8081", nil))
}

func main() {
	
}