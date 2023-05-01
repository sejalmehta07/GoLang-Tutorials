package main

import "fmt"

type Vertex1 struct {
	a int
	b int
}

func main() {
	var v1 = Vertex1{500, 600}
	var v2 = Vertex1{a: 100}
	var v3 = Vertex1{b: 200}
	var v4 = &Vertex1{1, 2}
	var v5 = Vertex1{}
	var v6 = &Vertex1{}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)

}
