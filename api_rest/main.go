package main

import ("fmt"
		"net/http"
		"log"
		"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true);
	router.HandleFunc("/health", Health);

	fmt.Println("Running server in http://localhost:2020");
	server := http.ListenAndServe(":2020", router) ;

	log.Fatal(server); // To catch problems
}

func Health(resp http.ResponseWriter, req *http.Request) {
	// TODO: Improve this health func
	fmt.Fprintf(resp, "Server is up and running...");
}