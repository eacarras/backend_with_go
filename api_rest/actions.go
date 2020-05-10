package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

// Dummy data
var persons_array = persons {
	person{"1a", "Aaron", "Carrasco", 22, 'S'},
	person{"2a", "Moises", "Carrasco", 28, 'C'},
	person{"1c", "Joel", "Castro", 32, 'V'},
}

func get_session() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost");

	if err != nil {
		panic(err)
	}

	return session
}

func Health(resp http.ResponseWriter, req *http.Request) {
	get_session() // Analyze if the connection is success
	fmt.Fprintf(resp, "Server is up and running...");
}

func PersonsList(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting list of persons...");

	var response []string;
	for i:= 0; i < len(persons_array); i ++ {
		response = append(response, persons_array[i].name);
	}
	json.NewEncoder(resp).Encode(response);
}

func PersonDetails(resp http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req);
	person_id := parms["id"];

	fmt.Printf("Getting the information of a person with id: %s...", person_id);

	for _, p := range persons_array {
		if p.Id == person_id {
			json.NewEncoder(resp).Encode(p);
		}
	}

	json.NewEncoder(resp).Encode(bad_request{"Not person found..", 400});
}

func PersonAdd(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body);

	var person_data person;
	err := decoder.Decode(&person_data);

	if err != nil {
		panic(err)
	}

	fmt.Printf("Adding a new person in the collection..");
	defer req.Body.Close();

	fmt.Println(person_data)
	persons_array = append(persons_array, person_data);
}
