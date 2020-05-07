package main

import "fmt"

func main() {
	var first_value float32;
	var second_value float32;
	
	const nameA string = "Aaron";
	const nameM string = "Moises";
	const nameC string = "Carlos";

	initial_message(nameA, nameM, nameC);
	
	fmt.Printf("Please write the first decimal value: ");
	fmt.Scanf("%f", &first_value);

	fmt.Printf("Please write the second decimal value: ");
	fmt.Scanf("%f", &second_value);

	operations(first_value, second_value);
}

func initial_message(names ...string) {
	/*
		This for and output can be improved
	*/
	fmt.Print("Welcome for the virtual calc ");
	
	for _, name := range names {
		fmt.Print(name + " ");
	}
	fmt.Print("\n")
}

func operations(first_value float32, second_value float32) {
	/*
		This func can be improve with a switch statement
	*/
	calc := func(fvalue float32, svalue float32, operator byte) float32{
		switch operator {
			case '+':
				return fvalue + svalue;
			case '-':
				return fvalue - svalue;
			case '*':
				return fvalue * svalue;
			case '/':
				return fvalue / svalue;
			default:
				return fvalue;
		}
	}

	fmt.Printf("The sum is: %.2f\n", calc(first_value, second_value, '+'));
	fmt.Printf("The difference is: %.2f\n", calc(first_value, second_value, '-'));
	fmt.Printf("The multiplication is: %.2f\n", calc(first_value, second_value, '*'));
	fmt.Printf("The division is: %.2f\n", calc(first_value, second_value, '/'));
}
