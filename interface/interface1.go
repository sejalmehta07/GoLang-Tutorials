package main

import "fmt"

type school interface {
	student() string
}
type subject struct {
	name string
}

func (s *subject) student() string {
	sub := subject{"SQL"}

	return &sub
}

type MyString string

func (ss MyString) student() string {
	greeting := "C"
	return greeting
}

func main() {
	var s *school            //interface
	str := subject{"hii"}    //struct
	ms := MyString("heeloo") //string type
	s = ms                   //MyString implements interface
	s = &str                 //*subject implements interface
	fmt.Println(s.student())
}
