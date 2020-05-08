package main

type person struct {
	id string `json:"identifier"`
	name string 
	last_name string `json:"lastName"`
	age int
	civil_state byte `json:"civilState"`
};

type bad_request struct {
	message string
	status int
}

type persons []person;
