package main

import (
	"fmt"
	"time"
)

func foo(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	message1 := "Learn GoLang"
	// message2 := "Learn Data Structures"
	go func() {
		fmt.Println(message1)
	}()
	// time.Sleep(10000 * time.Millisecond)
	go foo("goroutine")
	go func(message2 string) {
		fmt.Println(message2)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")

}
