package main

import ("fmt"
		"net/http"
		"log"
		"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true);
	// Routes
	router.HandleFunc("/health", Health);
	router.HandleFunc("/persons", PersonsList);
	router.HandleFunc("/persons/{id}", PersonDetails);

	fmt.Println("Running server in http://localhost:2020");
	server := http.ListenAndServe(":2020", router) ;

	log.Fatal(server); // To catch problems
}

func Health(resp http.ResponseWriter, req *http.Request) {
	// TODO: Improve this health func
	fmt.Fprintf(resp, "Server is up and running...");
}

func PersonsList(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting list of persons...");
	fmt.Fprintf(resp, "Persons: ");
}

func PersonDetails(resp http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req);
	person_id := parms["id"];

	fmt.Printf("Getting the information of a person with id: %s...", person_id);
	fmt.Fprintf(resp, "Person Detail: ");
}
