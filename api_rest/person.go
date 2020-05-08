package main

type person struct {
	Id string `json:"identifier"`
	name string 
	Last_name string `json:"lastName"`
	age int
	Civil_state byte `json:"civilState"`
};

type bad_request struct {
	message string
	status int
}

type persons []person;
