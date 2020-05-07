package main

import "fmt"

func main() {
	const last_index int = 5;
	movies := []string {"Terminator", "Avengers", "Lion King"};

	fmt.Println("For loop with index..");

	// One way to use for loop
	for i:= 1; i <= last_index; i ++ {
		fmt.Printf("The index is: %d\n", i);
	}

	// Like a forEach
	for _, movie := range movies {
		fmt.Println("For each example: " + movie);
	}
}