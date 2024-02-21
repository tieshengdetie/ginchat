package main

import "fmt"

func main() {
	ch1 := make(chan struct{})
	ch1 <- struct{}{}
	<-ch1
	fmt.Print(22)
}
