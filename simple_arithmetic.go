package main

// Add - Exported, a+b returns int
func Add(a, b int) int {
	// TODO: test with floats
	return a + b
}

// Subtract - Exported, a-b return int
func Subtract(a, b int) int {
	// TODO: test with floats
	return a - b
}

// Multiply - Exported, take a * b and return as int
func Multiply(a, b int) int {
	// TODO What about floats?
	return a * b
}

// Divide - Exported, take a / b and returns as int
func Divide(a, b int) int {
	// TODO: Turn this into a float and see how it affects tests
	return a / b
}
