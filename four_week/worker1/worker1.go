package main

import "fmt"

func fibonacci(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	var a, b uint64 = 0, 1
	var result uint64

	for i := 2; i <= n; i++ {
		result = a + b
		a, b = b, result
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}
