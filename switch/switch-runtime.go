package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOARCH; os {
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s", os)
	}
}
