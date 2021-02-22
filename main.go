package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Test struct {
	string string `json:"string`
	integer int `json:"integer"`
	bool bool `json:"bool"`
}

type Tests []Test

func TestingFunc(w http.ResponseWriter, r *http.Request){
	tester := Tests{
		Test{string: "This is a string", integer: 1, bool: true},
	}
	fmt.Println("EndPoint Hit!")
	json.NewEncoder(w).Encode(tester)
}

func landingPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Endpoint Hit!")
}

func handleRequest(){
	http.HandleFunc("/", landingPage)
	http.HandleFunc("/test", TestingFunc)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}