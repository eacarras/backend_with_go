package main

type person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Last_name string `json:"last_name"`
	Age int `json:"age"`
	Civil_state int `json:"civil_state"`
};

type bad_request struct {
	message string
	status int
}

type persons []person;
