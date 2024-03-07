package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	s := []int{1, 1, 1}
	fs(s)
	fmt.Println(s)
}
func fs(s []int) {
	s[0] = 22
}
