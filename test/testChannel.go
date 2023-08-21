package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("非main函数返回值为", mirroredQuery())
	time.Sleep(time.Second * 5)
	fmt.Println("main函数结束")
}

func mirroredQuery() string {
	responses := make(chan string)
	go func() {
		responses <- request("A", 3)
		fmt.Println("A协程结束")
	}()

	go func() {
		responses <- request("B", 1)
		fmt.Println("B协程结束")
	}()
	return <-responses
}

func request(hostname string, n int) (response string) {
	time.Sleep(time.Second * time.Duration(n))
	return hostname
}
