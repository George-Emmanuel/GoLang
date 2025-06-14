package main

import (
	"fmt"
)

func main() {

	//Declaring an empty slice in shorthand
	//Appening elements to the slice
	a := []int{}
	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
}
