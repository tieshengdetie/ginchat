package main

import "fmt"

func main() {
	arr := [...]int{12, 21, 55, 123, 888, 3, 5, 34, 11}
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
