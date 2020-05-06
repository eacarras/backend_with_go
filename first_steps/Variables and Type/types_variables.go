package main

import "fmt"

func main() {
	const sum int = 8 + 9;

	var name string = "Aaron"
	second_string := fmt.Sprintf(" the sum is: %d\n", sum)

	fmt.Printf("Hello " + name + second_string);
}