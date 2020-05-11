package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	router := NewRouter();

	fmt.Println("Running server in http://localhost:2020");
	server := http.ListenAndServe(":2020", router) ;

	log.Fatal(server); // To catch problems
}
