// Package addition provides simple arithmetic operations
package addition

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// AddMultiple returns the sum of multiple integers
func AddMultiple(numbers ...int) int { //numbers is treated as a slice of integers ([]int)
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
