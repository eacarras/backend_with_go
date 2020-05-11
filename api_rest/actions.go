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

func response(resp http.ResponseWriter, status int, result person) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	json.NewEncoder(resp).Encode(result);
}

func response_array(resp http.ResponseWriter, status int, result []person) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	json.NewEncoder(resp).Encode(result);
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

	response_array(resp, 200, results)
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

	response(resp, 200, result)
}

func PersonUpdate(resp http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req);
	person_id := parms["id"];

	fmt.Printf("Getting the information of a person with id: %s...\n", person_id);

	if !bson.IsObjectIdHex(person_id) {
		resp.WriteHeader(404)
		return
	}

	decoded_id := bson.ObjectIdHex(person_id)
	decoder := json.NewDecoder(req.Body)
	
	var person_data person
	err_decoder := decoder.Decode(&person_data)

	if err_decoder != nil {
		panic(err_decoder)
		resp.WriteHeader(500)
		return
	}
	defer req.Body.Close()

	document := bson.M{"_id": decoded_id}
	change := bson.M{"$set": person_data}
	err := collection.Update(document, change)

	if err != nil {
		panic(err)
		resp.WriteHeader(404)
		return
	}

	response(resp, 200, person_data)
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
	response(resp, 200, person_data)
}
