package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	rollNumber int `required max:100`
	name       string
	class      string
	subjects   []string
	marks      []int
}

func main() {
	// sejal := Student{
	// 	rollNumber: 21735,
	// 	name:       "Sejal Mehta",
	// 	class:      "MCA-II",
	// 	marks: []int{
	// 		10, 20, 30,
	// 	},
	// 	subjects: []string{
	// 		"OOPS", "Java", "SQL",
	// 	},
	// }
	// // fmt.Println(sejal)
	// var harshu Student
	// harshu = sejal
	// fmt.Println(harshu.name)

	st := reflect.TypeOf(Student{})
	field, _ := st.FieldByName("rollNumber")
	fmt.Println(field)

}
