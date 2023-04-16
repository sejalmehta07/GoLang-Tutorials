package main

import "fmt" //fmt package includes i/o functions
func main() {
	//////types of values- boolean, string, integer, float

	/*fmt.Printf(false && true)          (error aa rha hai && cannot use false && true (untyped bool constant false)
	  as string value in argument to fmt.PRINTFCOMPILERINCOMPATIBLEASSIGN)*/
	fmt.Println(!true)
	fmt.Println(false || true)
	fmt.Println("100+200=", 100+200)
	fmt.Println(10 + 20)
	fmt.Println("this is" + " a string")
	fmt.Println(13.3526 + 536.2571)

}
