package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func Health(resp http.ResponseWriter, req *http.Request) {
	// TODO: Improve this health func
	fmt.Fprintf(resp, "Server is up and running...");
}

func PersonsList(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting list of persons...");

	// Dummy data
	persons := persons {
		person{"1a", "Aaron", "Carrasco", 22, 'S'},
		person{"2a", "Moises", "Carrasco", 28, 'C'},
		person{"1c", "Joel", "Castro", 32, 'V'},
	}

	var response []string;
	for i:= 0; i < len(persons); i ++ {
		response = append(response, persons[i].name);
	}
	json.NewEncoder(resp).Encode(response);
}

func PersonDetails(resp http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req);
	person_id := parms["id"];

	fmt.Printf("Getting the information of a person with id: %s...", person_id);
	// Dummy data
	persons := persons {
		person{"1a", "Aaron", "Carrasco", 22, 'S'},
		person{"2a", "Moises", "Carrasco", 28, 'C'},
		person{"1c", "Joel", "Castro", 32, 'V'},
	}

	for _, p := range persons {
		if p.Id == person_id {
			json.NewEncoder(resp).Encode(p);
		}
	}

	json.NewEncoder(resp).Encode(bad_request{"Not person found..", 400});
}
