package main

import "fmt"

func main() {
	// annonymous function- inside main function  & with no name
	// for i := 0; i < 10; i++ {
	// 	func(i int) {
	// 		if i == 5 {
	// 			fmt.Println("hii, this is annonymous function")
	// 			fmt.Println(i)
	// 		}
	// 	}(i) // - these closing parantheses- the function gets immediately invoked due to this
	// }
	j := 1
	f := func() {
		fmt.Println(j)
	}
	f()

}
