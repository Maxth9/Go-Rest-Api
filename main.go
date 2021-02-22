package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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

func TestingPostFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Post Endpoint Hit!")
}

func landingPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Endpoint Hit!")
}

func handleRequest(){
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", landingPage)
	r.HandleFunc("/test", TestingFunc).Method("Get")
	r.HandleFunc("/test", TestingPostFunc).Method("POST")
	r.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	handleRequest()
}