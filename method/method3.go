package main

import "fmt"

type MyString string

// declaring without struct type ( here, string type with a method)
func (ms MyString) fun() {
	fmt.Println(ms)
}
func main() {
	ms := MyString("hii")
	ms.fun()
}
