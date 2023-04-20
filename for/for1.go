package main

import (
	"fmt"
	"math"
)

func main() {
	//basic for loop
	//3 components - init, condition, post
	var i float64
	for i = 1; i <= 10; i++ {
		fmt.Println(math.Sqrt(i))
	}

	// for loop without init and post
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}
	// for loop with break statement and without condition
	for {
		fmt.Println("For with Break")
		break
	}
	// infinite loop
	// for {
	// 	fmt.Println("Inside for")

	// }

}
