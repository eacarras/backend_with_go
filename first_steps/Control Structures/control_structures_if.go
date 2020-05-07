package main

import ("fmt"
		"os"
		"strconv"
	)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("You don't send your name and your age");
	} else {
		fmt.Printf("Welcome %s! You have %s years old and you are\n", os.Args[1], os.Args[2]);
		age, err := strconv.Atoi(os.Args[2]);

		if err != nil { 
			fmt.Printf("You don't send a number !\n");
			os.Exit(1); 
		} 

		if age < 18 {
			fmt.Println("MINOR AGE!");
		} else {
			fmt.Println("OLD !! :)");
		}
	}
}