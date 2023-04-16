package main

import "fmt"

const name = "sejal" //package level
func main() {
	const c = 10 // function level
	fmt.Println(c)
	fmt.Println(name)
	fmt.Println(float64(c))

}
