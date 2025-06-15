package main

import (
	"errors"
	"fmt"
)

func showErrors() {
	// Example of creating and using a custom error
	err := errors.New("this is a custom error")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example of wrapping an error
	wrappedErr := fmt.Errorf("printing wrapped error ==> %w", err)
	if wrappedErr != nil {
		fmt.Println("Wrapped Error:", wrappedErr)
	}

	// Example of checking for a specific error
	if errors.Is(wrappedErr, err) {
		fmt.Println("The wrapped error matches the original error.")
	}

	//Division by zero error example
	divideByZero := func() error {
		return errors.New("division by zero error")
	}
	if err := divideByZero(); err != nil {
		fmt.Println("Error:", err)
	}
}

/* Uncomment the following lines to run the showErrors function
func main() {
	showErrors()
}*/
