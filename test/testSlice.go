package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	slice2 := slice1[1:5]
	cutSlice(slice2)
	fmt.Println(len(slice1))
	fmt.Println(slice2)
}
func cutSlice(s []int) {
	s[0] = 1
}
