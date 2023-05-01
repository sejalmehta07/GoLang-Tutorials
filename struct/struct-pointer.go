package main

import "fmt"

// type Vertexpq struct {
// 	p int
// 	q int
// }

type student struct {
	name   string
	age    int
	result bool
}

func newStudent(n string) *student {
	s := student{name: n}
	s.age = 22
	s.result = true
	return &s
}

func main() {
	// v := Vertexpq{10, 20}
	// s := &v 			////////accessing struct using pointer////////////
	// s.p = 100
	// fmt.Println(s)
	// fmt.Println(s.p)
	st := student{"tom", 10, false}
	fmt.Println(st)
	st.age = 100
	fmt.Println(st)
	ns := newStudent("harry")
	fmt.Println(ns)

}
