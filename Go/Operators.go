package main

import (
	"fmt"
)

func main() {
	num1 := 10
	num2 := 20

	// Arithmetic Operators
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	diff := num2 - num1
	fmt.Println("Difference:", diff)

	product := num1 * num2
	fmt.Println("Product:", product)

	quotient := num2 / num1
	fmt.Println("Quotient:", quotient)

	rem := num2 % num1
	fmt.Println("Remainder:", rem)

	//Relational Operators
	isEqual := num1 == num2
	fmt.Println("Is Equal:", isEqual)

	isNotEqual := num1 != num2
	fmt.Println("Is Not Equal:", isNotEqual)

	isGreater := num2 > num1
	fmt.Println("Is Greater:", isGreater)

	isLess := num1 < num2
	fmt.Println("Is Less:", isLess)

	isGreaterOrEqual := num2 >= num1
	fmt.Println("Is Greater or Equal:", isGreaterOrEqual)

	isLessOrEqual := num1 <= num2
	fmt.Println("Is Less or Equal:", isLessOrEqual)

	// Logical Operators
	isTrue := true
	isFalse := false
	andResult := isTrue && isFalse
	fmt.Println("Logical AND:", andResult)
	orResult := isTrue || isFalse
	fmt.Println("Logical OR:", orResult)
	notResult := !isTrue
	fmt.Println("Logical NOT:", notResult)

	// Bitwise Operators
	bitwiseAnd := num1 & num2
	fmt.Println("Bitwise AND:", bitwiseAnd)
	bitwiseOr := num1 | num2
	fmt.Println("Bitwise OR:", bitwiseOr)
	bitwiseXor := num1 ^ num2
	fmt.Println("Bitwise XOR:", bitwiseXor)
	bitwiseNot := ^num1
	fmt.Println("Bitwise NOT:", bitwiseNot)
	bitwiseLeftShift := num1 << 1
	fmt.Println("Bitwise Left Shift:", bitwiseLeftShift)
	bitwiseRightShift := num2 >> 1
	fmt.Println("Bitwise Right Shift:", bitwiseRightShift)

	// Assignment Operators
	num3 := 30
	num3 += num1
	fmt.Println("After += :", num3)
	num3 -= num2
	fmt.Println("After -= :", num3)
	num3 *= num1
	fmt.Println("After *= :", num3)
	num3 /= num2
	fmt.Println("After /= :", num3)
	num3 %= num1
	fmt.Println("After %= :", num3)
	num3 <<= 1
	fmt.Println("After <<= :", num3)
	num3 >>= 1
	fmt.Println("After >>= :", num3)
	num3 &= num1
	fmt.Println("After &= :", num3)
	num3 |= num2
	fmt.Println("After |= :", num3)
	num3 ^= num1
	fmt.Println("After ^= :", num3)
	num3 = 100 // Reset num3 for clarity
	num3 = 100
	fmt.Println("Reset num3 to:", num3)
}
