package main

import "fmt"

func main() {
	var movies [3]string;
	var matrix [2][2]int;

	names := [3]string{"Aaron", "Enrique", "Joel"}
	slice := []int{1, 2, 3, 4}
	
	var index int;
	var new_value int;

	matrix[0][0] = 1
	matrix[1][1] = 1
	matrix[1][0] = 0
	matrix[0][1] = 0

	movies[0] = "Avengers"
	movies[1] = "Avengers: Age of Ultron"
	movies[2] = "Avengers: Infinity War"

	fmt.Printf("Array of one dimension examples:\nMovies array: %v\nNames array %v\n", movies, names);
	fmt.Printf("\nArray of multiple dimension examples:\n%v \n", matrix);

	show_slice_and_length(slice);
	fmt.Printf("Please write a index less than the length of the slice: ")
	fmt.Scanf("%d", &index);
	if (index <= len(slice)) {
		fmt.Printf("Your slice is: %v\n", slice[:index]);
	} else {
		fmt.Printf("You send a bad index ..\n");
	}
	
	fmt.Printf("Please write a new value to be inserted: ");
	fmt.Scanf("%d", &new_value);
	slice = append(slice, new_value);
	show_slice_and_length(slice);
}

func show_slice_and_length(slice []int) {
	fmt.Printf("\nSlice example:\nSlice: %v\nLength: %d\n", slice, len(slice));
}
