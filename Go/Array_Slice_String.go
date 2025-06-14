package main

import (
	"fmt"
	"strings"
)

func main() {

	//Declaring	and manipulating an array
	var arr [5]int
	arr[0] = 6
	arr[1] = 7
	arr[2] = 8
	arr[3] = 9
	arr[4] = 0
	fmt.Println("Array ==> ", arr)

	//Declaring an empty slice in shorthand
	//Appending elements to the slice
	a := []int{}
	a = append(a, 2, 3, 4, 5)
	fmt.Println("Slice ==> ", a)

	//Declaring String and splitting it
	//using Split function
	st := "Test"
	mst := "Test@gmail.com"
	split_string := strings.Split(mst, "@")
	fmt.Println("String ==> ", st)
	fmt.Println("Split_string ==> ", split_string)

	//Analysing Split String Array
	fmt.Println("Split_string First Value ==> ", split_string[0])
	fmt.Println("Split_string Last Value ==> ", split_string[len(split_string)-1])
	fmt.Println("Split_string Length ==> ", len(split_string))
	fmt.Println("Split_string Contains 'gmail' ==> ", strings.Contains(split_string[1], "gmail"))
	fmt.Println("Split_string Contains 'yahoo' ==> ", strings.Contains(split_string[1], "yahoo"))
}
