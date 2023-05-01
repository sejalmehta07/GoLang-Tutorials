package main

import "fmt"

type student struct {
	name   string
	age    int
	result bool
}

func (s student) subject(a, b, c string) {  // method subject
	fmt.Println("Subject 1 : ", a)
	fmt.Println("Subject 2 : ", b)
	fmt.Println("Subject 3 : ", c)
}
func (s student) displayDetails() {  // method displayDetails
	fmt.Println("Name : ", s.name)
	fmt.Println("Age : ", s.age)
	fmt.Println("Result : ", s.result)
}
func main() {
	s := student{
		name:   "Ashika",
		age:    20,
		result: true,
	}
	s.displayDetails()
	s.subject("Java", "C++", "SQL")
}
