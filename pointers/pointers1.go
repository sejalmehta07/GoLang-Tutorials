package main

import "fmt"

func value(n int) {
	n = 100
}
func addresss(ptr *int) {
	*ptr = 500
}
func main() {
	// i := 10
	// p := &i

	// fmt.Println(i)
	// fmt.Println(p)

	// *p = 20
	// fmt.Println(i)
	// fmt.Println(p)

	// var v1 int
	// v1 = 80
	// fmt.Printf("value of v1-- ")
	// fmt.Println(v1)
	// var p1 *int
	// p1 = &v1
	// *p1 = 40
	// fmt.Println(p1)
	// fmt.Printf("value of v1 again-- ")
	// fmt.Println(v1)

	// var a int = 10
	// fmt.Println(a)

	// var p *int
	// p = &a
	// fmt.Println(*p, a)
	// a = 20
	// fmt.Println(*p, a)

	n := 10
	fmt.Println("initial value : ", n)
	fmt.Println("address of n: ", &n)

	value(n)
	fmt.Println("after calling value function: ", n)
	addresss(&n)
	fmt.Println("after calling value addresss: ", n)

}
