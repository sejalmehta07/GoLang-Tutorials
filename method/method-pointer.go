package main

import "fmt"

type rectangle struct {
	l int // l- length
	b int // b- breadth
}

func (rec *rectangle) updateValues() {
	rec.l = 15								
	rec.b = 5
}
func (rec *rectangle) area() int {
	return rec.l * rec.b	
}
func main() {
	r := rectangle{10, 20}
	r.updateValues()
	a := r.area()
	fmt.Println(a)
}
