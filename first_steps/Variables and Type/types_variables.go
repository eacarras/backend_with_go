package main

import "fmt"

type struct_example struct {
	nombre string
	edad int
	sueldo float32
	muerto bool
}

func main() {
	const sum int = 8 + 9;
	var aaron = struct_example {
		"Aaron Carrasco",
		22,
		123.78,
		false,
	}
	var name string = "Aaron"

	second_string := fmt.Sprintf(" the sum is: %d\n", sum)
	struct_string := fmt.Sprintf("%v", aaron); // We can use %#v to add name fields

	fmt.Printf("Hello " + name + second_string);
	fmt.Println("This is a struct: " + struct_string); // We can improve this struct_string with loops
}
