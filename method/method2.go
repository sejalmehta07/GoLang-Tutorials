package main

import "fmt"

type person struct {
	name string
	id   int
}

func displayDetails(p person) { // displayDetails is a function here
	fmt.Println(p.id, " : ", p.name)
}
func main() {
	p := person{"tommy", 1001}
	displayDetails(p)

}
