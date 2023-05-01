package main

import "fmt"

func main() {
	m := make(map[int]string)
	m = map[int]string{
		1: "sejal",
		2: "mehta",
		3: "harshita",
	}

	m2 := map[string]int{
		"aa": 10,
		"bb": 20,
		"cc": 30,
	}
	m[4] = "mehta"
	m3 := m2
	fmt.Println(m3)
	fmt.Println(m2)
	delete(m2, "cc")
	fmt.Println(m3)
	fmt.Println(m2)

}
