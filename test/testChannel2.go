package main

import "fmt"

func main() {
	go func() {

	}()
	ch1 := make(chan bool)
	<-ch1
	fmt.Print(22)
}
