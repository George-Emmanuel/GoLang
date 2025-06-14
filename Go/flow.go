package main

import (
	"fmt"
)

func main() {
	// For loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// For loop with range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Nested for loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("Outer loop %d, Inner loop %d\n", i, j)
		}
	}

	// Foreach loop (using for range)
	for _, value := range numbers {
		fmt.Println("Foreach loop value:", value)
	}

	// While loop (using for)
	i := 0
	for i < 5 {
		fmt.Println("While loop iteration:", i)
		i++
	}

	// Do-while loop (using for)
	i = 0
	for {
		fmt.Println("Do-while loop iteration:", i)
		i++
		if i >= 5 {
			break
		}
	}

	// Switch statement
	switch day := 3; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	}

	// If-else statement
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// If-else if-else statement
	if y := 10; y > 10 {
		fmt.Println("y is greater than 10")
	} else if y == 10 {
		fmt.Println("y is equal to 10")
	} else {
		fmt.Println("y is less than 10")
	}

	// Infinite loop
	for {
		fmt.Println("This is an infinite loop. Press Ctrl+C to stop.")
		break // Remove this line to make it truly infinite
	}

	// Break statement
	for j := 0; j < 10; j++ {
		if j == 5 {
			fmt.Println("Breaking the loop at j =", j)
			break
		}
		fmt.Println("j =", j)
	}

	// Continue statement
	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			fmt.Println("Skipping even number:", k)
			continue
		}
		fmt.Println("Odd number:", k)
	}

	// Fallthrough in switch
	switch num := 2; num {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
		fallthrough
	case 3:
		fmt.Println("Number is 3 or greater")
	default:
		fmt.Println("Number is not 1, 2, or 3")
	}

	// Defer statement
	defer fmt.Println("This will be printed at the end of the main function.")
	fmt.Println("End of main function execution.")
}
