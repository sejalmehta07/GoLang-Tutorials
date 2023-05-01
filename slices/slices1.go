package main

import "fmt"

func main() {
	/// SLICE - New Datatype in go

	// sl := make([]int, 5) // make- keyword for declaring slices
	// fmt.Println(sl)
	// fmt.Println(len(sl))
	// i := 0
	// for i < len(sl) {
	// 	sl[i] = i
	// 	i++
	// }
	// fmt.Println(sl)

	// newSlice := make([]int, 5)
	// newSlice[0] = 1
	// newSlice[1] = 2
	// newSlice[2] = 3
	// fmt.Println(newSlice)

	// a := make([]int, 3)

	// a = append(a, 1)
	// a = append(a, 2)
	// a = append(a, 3)

	// b := make([]int, 3)

	// b = append(b, 8)
	// b = append(b, 9)
	// b = append(b, 10)

	// fmt.Println(a)
	// fmt.Println(b)

	// c := append(a, b...)
	// fmt.Println(c)

	// temp := []int{10, 20, 30}
	// temp = append(temp, 40, 50)

	// flag := []int{60}
	// flag = append(temp, flag...)
	// fmt.Println(flag)

	///// copy function - copies one slice into another

	// original := []int{100, 200, 300, 199, 822, 152} // declaration and initialization in the samw line

	// copied := make([]int, len(original))
	// copy(copied, original)

	// u := original[2:4]
	// fmt.Println(u)

	//// two dimensional slices

	// twoDim := make([][]int, 3)
	// for i := 1; i < 3; i++{}

	sli := [3]string{"aa", "bb", "cc"}
	for _, v := range sli {
		fmt.Println(v)
	}

}
