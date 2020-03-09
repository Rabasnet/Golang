//this has to be attached to the database for the actual scenario

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	ID        string   `json:"id, omitempty"`
	FirstName string   `json: "firstname, omitempty"`
	LastName  string   `json: "lastname, omitempty"`
	Address   *Address `json:"address, omitempty"` //here Address is of type Address struct
}

//Address is a struct that holds all types of info
type Address struct {
	city  string `json:"city, omitempty"`
	State string `json: "state, omitempty"`
}

var people []person

func main() {
	myRouter := mux.NewRouter()
	people = append(people, person{ID: "11", FirstName: "rabi", LastName: "basnet", Address: &Address{city: "irving", State: "Texas"}})
	people = append(people, person{ID: "12", FirstName: "rajina", LastName: "bista", Address: &Address{city: "bhaktapur", State: "Nepal"}})
	people = append(people, person{ID: "13", FirstName: "binita", LastName: "thapa"})

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/people/{id}", getHandler).Methods("GET")
	myRouter.HandleFunc("/people/{id}", postHandler).Methods("POST")
	myRouter.HandleFunc("/people/{id}", deleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8881", myRouter))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my first api")
}
func getHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&person{}) //encoder for get method
}

//this is for post or create
//people is slice, Pperson is the new item here, person is the struct,
func postHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var Pperson person
	Pperson.ID = params["id"]
	_ = json.NewDecoder(r.Body).Decode(&Pperson)
	people = append(people, Pperson)
	json.NewEncoder(w).Encode(people) //this adds and then prints the whole slice // decoder and then encoder for post method

}

//delete
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...) //remember that append takes ...
			break
		}

	}
	json.NewEncoder(w).Encode(people) // encoder for the delete method
}
