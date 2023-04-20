package main

import "fmt"

func main() {
	// if else
	num := 100
	if num > 50 {
		fmt.Println("Inside if block")
	} else {
		fmt.Println("Inside else block")
	}

	// if-else-if
	if varr := 200; varr < 50 {
		fmt.Println("Inside if block of if-else-if")
	} else if varr > 51 && varr < 100 {
		fmt.Println("inside else if block")
	} else {
		fmt.Println("inside else block")
	}
}
