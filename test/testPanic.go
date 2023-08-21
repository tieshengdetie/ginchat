package main

import "fmt"

func main() {

	fmt.Println("开始时执行程序")

	defer func() {
		if data := recover(); data != nil {
			fmt.Println(data)
			fmt.Println("程序恢复喽")
		}
	}()

	panic("程序panic了")

}
