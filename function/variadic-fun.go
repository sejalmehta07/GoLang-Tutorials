package main

import "fmt"

func sum(n ...int) {
	fmt.Println(n)
	total := 0
	for _, val := range n {
		total += val
	}
	fmt.Println(total)
}
func main() {
	sum(1, 2, 3, 4, 5)
	sum(1, 2, 3)
	n := []int{1, 2, 3, 4}
	sum(n...)
}
