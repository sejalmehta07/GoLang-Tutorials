package main

import (
	"fmt"
	"time"
)

func main() {
	// Now function under the time package tells the current time
	t := time.Now()
	fmt.Println(t)
	s := time.Now().Weekday()
	fmt.Println(s)
	k := time.Now().Day()
	fmt.Println(k)
	p := time.Now().Month()
	fmt.Println(p)
	m := time.Now().Hour()
	fmt.Println(m)
	h := time.Now().Weekday()
	fmt.Println(h)

	//switch
	day := 5
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
	whatAmI := func(xy interface{}) {
		switch par := xy.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("Don't know type %T\n", par)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("my name is sejal mehta")

}
