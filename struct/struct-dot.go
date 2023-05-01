package main

import "fmt"

type Vertex struct {
	a int
	b int
}

func main() {
	v := Vertex{2, 4}
	v.a = v.b * 10   /// accessing fields of a struct by dot operator
	fmt.Println(v.a)

}
