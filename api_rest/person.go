package main

type person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Last_name string `json:"last_name"`
	Age int `json:"age"`
	Civil_state int `json:"civil_state"`
};

type message struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

type persons []person;
