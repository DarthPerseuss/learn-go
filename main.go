package main

import (
	"fmt"
	"learn-go/calc"
	"learn-go/strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(calc.Add(1, 2), "Addition")
	fmt.Println(calc.Subtract(1, 2), "Subtraction")
	fmt.Println(calc.Multiply(1, 2), "Multiplication")
	fmt.Println(calc.Divide(1, 2), "Division")
	fmt.Println(calc.Factorial(5), "Factorial")

	fmt.Println("\nTesting string funcs")
	fmt.Println("\n", strings.Reverse("Dynamite"), "Reveresed string")
	fmt.Println(strings.IsPalindrome("madam"), "madam is palindrome")
	fmt.Println(strings.IsPalindrome("maddam"), "maddam is palindrome")
	fmt.Println(strings.IsPalindrome("Dynamite"), "Dynamite is not palindrome")
}
