package main

import "fmt"

func main() {
	fmt.Println("start")
	a, b := 1, 0
	c := a / b
	defer fmt.Println(c)

	fmt.Println("end")
}
