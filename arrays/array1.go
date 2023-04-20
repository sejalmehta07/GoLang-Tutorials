package main

import (
	"fmt"
)

func main() {
	//one-dimensional arrays:-------
	var arr [5]string
	fmt.Println(arr)
	arr[2] = "hi" // setting elements at a particular index
	arr[3] = "sejal"
	fmt.Println(arr)
	fmt.Println(arr[3])
	arr2 := [3]float64{1.353, 3.363, 9.124} // declaration and initialization
	fmt.Println(arr2)

	//2-D array:-----
	var new2Darrray [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			new2Darrray[i][j] = i + j
		}
	}
	fmt.Println(new2Darrray)

}
