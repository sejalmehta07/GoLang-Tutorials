package main

import "fmt"

// func sum(a int, b int) int {
// 	sum := a + b
// 	sub := a - b
// 	return sum, sub
// }
// func concat(a, b string) (string, string) {
// 	return a, b
// }

func returnMulVal(a, b int) (int, error) {
	if a < 15 {
		return a, fmt.Errorf("a is less than 15")
	}
	return 10, nil
}
func main() {
	// sum, sub := sum(2, 4)
	// fmt.Println(sum)
	// fmt.Println(sub)
	// fmt.Println(concat("sejal ", " mehta"))
	// _, str := concat("HELLO ", "WORLD!")
	// fmt.Println(str)
	var a, b = returnMulVal(10, 20)
	fmt.Println(a)
	fmt.Println(b)
	// ////// only subset of the returned values
	// _, subset := returnMulVal()
	// fmt.Println(subset)
}
