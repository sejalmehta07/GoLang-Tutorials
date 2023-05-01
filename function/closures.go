package main

import "fmt"

func outerfunc() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}
func main() {
	innerfunc := outerfunc()

	fmt.Println(innerfunc())
	fmt.Println(innerfunc())
	fmt.Println(innerfunc())
	fmt.Println(innerfunc())
	fmt.Println(innerfunc())
	//new case to identify the uniqueness
	innerfunction := outerfunc()
	fmt.Println(innerfunction())

}
