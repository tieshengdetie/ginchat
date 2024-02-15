package main

import "fmt"

func main() {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
			chan2 <- i
		}
	}()
	defer close(chan1)
	defer close(chan2)
	for {
		select {
		case x := <-chan1:
			fmt.Println(x)
		case y := <-chan2:
			fmt.Println(y)
			//default:
			//	fmt.Println("当前channel里无值")
		}

	}

}
