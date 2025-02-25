package main

import (
	"fmt"
)

// First option is to run a for loop and sum up the values
// Time complexity of this solution is O(n)
func sum_to_n_a(n int) int {
	var sum int
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum

}

// Second option is to use the formula n(n+1)/2
// Time complexity of this solution is O(1)
func sum_to_n_b(n int) int {
	sum := n * (n + 1) / 2
	return sum
}

// Third option is to use recursion
// Time complexity of this solution is O(n)
func sum_to_n_c(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum_to_n_c(n-1)
}

func main() {
	fmt.Println(sum_to_n_a(10))
	fmt.Println(sum_to_n_b(10))
	fmt.Println(sum_to_n_c(10))
}
