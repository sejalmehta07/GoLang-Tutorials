package main

import (
	"fmt"
	"time"
)

func display(greet string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(greet)
}
func main() {
	go display("hi harshu")
	time.Sleep(100 * time.Millisecond)
	display("hii sejal")
}
