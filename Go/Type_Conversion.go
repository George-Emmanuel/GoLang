package main

import (
	"fmt"
	"strconv"
)

func type_conversion() {

	//Integer to Float Conversion
	var i1 int = 42
	var f1 = float64(i1)
	fmt.Println("Integer to Float64 Conversion: ", f1)

	//Float to Integer Conversion
	var f2 float64 = 42
	var i2 = int(f2)
	fmt.Println("Float to Integer Conversion: ", i2)

	//String to Integer Conversion
	var s = "12345"
	var i3, err = strconv.Atoi(s)
	// Note: Strconv.Atoi returns two values, the converted integer and an error.
	fmt.Println("String to Integer Conversion: ", i3)
	fmt.Println("String to Integer Conversion ERROR: ", err)

	//Integer to String Conversion
	var i4 = 42
	var s1 = strconv.Itoa(i4)
	fmt.Println("Integer to String Conversion: ", s1)
	fmt.Println("Integer to String Conversion ERROR: ", err)
}

func main() {
	type_conversion()
}
