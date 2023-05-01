package main

import "fmt"

func main() {
	// num := []int{1, 2, 3}
	// sum := 0
	// for _, items := range num { //////// a variable can also be used Instead of _ , that will represent the index
	// 	sum += items
	// }
	// fmt.Println(sum)

	// //// if we need to get the index:--

	// for i, item := range num {
	// 	fmt.Println(i, item)
	// }

	map1 := map[int]string{1: "diya", 2: "palak", 3: "mehak", 4: "pooja"}
	for k, val := range map1 {
		fmt.Printf("%v -> %s\n", k, val)
	}

	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }
	fmt.Println("Keys of map1 :--")
	for k := range map1 {
		fmt.Printf("%v\t", k)
	}
	fmt.Println("")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
