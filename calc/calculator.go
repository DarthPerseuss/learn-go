package calc

func Add(x int, y int) int {
	return x + y
}

func Subtract(x int, y int) int {
	return x - y
}

func Multiply(x int, y int) int {
	return x * y
}

func Divide(x float64, y float64) float64 {
	return x / y
}

func Factorial(x int) int {
	// This calculates the factorial of an integer and also
	// handles negative values by returning an error type.

	if x < 0 {
		return -1
	}
	product := 1
	for i := 1; i <= x; i++ {
		product *= i
	}
	return product
}
