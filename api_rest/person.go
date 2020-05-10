package main

type person struct {
	Id string
	Name string 
	Last_name string
	Age int
	Civil_state int
};

type bad_request struct {
	message string
	status int
}

type persons []person;
