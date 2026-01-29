package main

import (
	"fmt"
)

func ds_types() {
	n := 10

	// Pointer to n.
	// Here, we create a pointer to the variable n.
	// and assign it to the variable add.
	// The pointer add now holds the address of n
	add := &n
	val := *add
	fmt.Println("The value of n is:", n)
	fmt.Println("The address of n is:", add)
	fmt.Println("The value of add is:", val)

	//Structs
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person Age:", p.Age)

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	// Appending to a slice
	slice = append(slice, 6)
	fmt.Println("Slice after appending:", slice)

	// Maps
	mapExample := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("Map Example:", mapExample)
	// Accessing map values
	fmt.Println("Value for 'two':", mapExample["two"])
	// Adding a new key-value pair to the map
	mapExample["four"] = 4
	fmt.Println("Map after adding 'four':", mapExample)
	// Deleting a key-value pair from the map
	delete(mapExample, "one")
	fmt.Println("Map after deleting 'one':", mapExample)
}

/* Uncomment the following line to run the ds_types function
func main() {
	ds_types()
}
*/
