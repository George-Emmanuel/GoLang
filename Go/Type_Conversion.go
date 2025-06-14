package main

import (
	"fmt"
	"strconv"
)

// ***** Printf function is used to format and print output in Go.
// ***** strconv package provides functions for converting between strings and other types.
// ***** Println function is used to print output in Go without formatting.

func type_conversion() {

	//Integer to Float Conversion
	var i1 int = 3
	f1 := float64(i1)
	fmt.Printf("Integer to Float64 Conversion: %.1f\n", f1)

	//Float to Integer Conversion
	var f2 float64 = 3.12
	i2 := int(f2)
	fmt.Printf("Float to Integer Conversion: %v\n", i2)

	//String to Integer Conversion
	var s = "12345"
	var i3, err = strconv.Atoi(s)
	// Note: Strconv.Atoi returns two values, the converted integer and an error.
	fmt.Printf("String to Integer Conversion: %v\n", i3)
	fmt.Printf("String to Integer Conversion ERROR: %v\n", err)

	//Integer to String Conversion
	var i4 = 42
	var s1 = strconv.Itoa(i4)
	fmt.Printf("Integer to String Conversion: %s\n", s1)
	fmt.Printf("Integer to String Conversion ERROR: %v\n", err)
}

/* Uncomment to run the code
func main() {
	type_conversion()
}*/
