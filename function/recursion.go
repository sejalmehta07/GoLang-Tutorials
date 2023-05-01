package main

import "fmt"

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
func main() {
	f := factorial(3)
	fmt.Println(f)

	var fib func(k int) int
	fib = func(k int) int {
		if k < 2 {
			return k
		}
		return (k - 1) + (k - 2)
	}
	fmt.Println(fib(5))
}
