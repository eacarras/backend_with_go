package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func get_session() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost");

	if err != nil {
		panic(err)
	}
	return session
}

var collection = get_session().DB("go_test_db").C("persons")

func Health(resp http.ResponseWriter, req *http.Request) {
	get_session() // Analyze if the connection is success
	fmt.Fprintf(resp, "Server is up and running...");
}

func PersonsList(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting list of persons...");

	var results []person
	err := collection.Find(nil).Sort("_id").All(&results)
	
	if err != nil {
		resp.WriteHeader(500)
		log.Fatal(err)
		
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(results);
}

func PersonDetails(resp http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req);
	person_id := parms["id"];

	fmt.Printf("Getting the information of a person with id: %s...\n", person_id);

	if !bson.IsObjectIdHex(person_id) {
		resp.WriteHeader(404)
		return
	}

	decoded_id := bson.ObjectIdHex(person_id)
	
	result := person{}
	err := collection.FindId(decoded_id).One(&result)

	if err != nil {
		log.Fatal(err)
		resp.WriteHeader(404)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(result)
}

func PersonAdd(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body);

	var person_data person;
	err := decoder.Decode(&person_data);

	if err != nil {
		panic(err)
	}
	defer req.Body.Close(); // Close the body read

	fmt.Printf("Adding a new person in the collection..\n");
	
	err_insertion := collection.Insert(person_data)
	if err_insertion != nil {
		resp.WriteHeader(500)
		return	
	}

	fmt.Printf("Person added successfully..\n")
	
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(person_data)
}
