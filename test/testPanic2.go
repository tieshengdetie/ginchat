package main

import (
	"log"
)

func main() {
	go func() {
		//time.Sleep(2 * time.Second)
		panic("煎鱼焦了")
	}()

	log.Println("Go 语言编程之旅：一起用 Go 做项目")
}
