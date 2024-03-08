package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 存放文章信息的 Post 结构体
type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

func main() {

	// 初始化一个字典类型变量
	var books map[int]*Book = make(map[int]*Book)
	book1 := Book{Id: 1, Title: "Go Web 编程", Summary: "Go Web 编程入门指南", Author: "学院君"}
	books[book1.Id] = &book1
	// 通过 JSON 序列化字典数据
	data, _ := json.Marshal(books)
	file, err := os.OpenFile("/project/logs/logfile.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		//打印异常信息
		fmt.Println("open file err", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	//count := 100

	//for i := 0; i < count; i++ {
	//	logMsg := make(map[int]interface{})
	//	logMsg[i] = "这是一条日志消息"
	//	strObj, _ := json.Marshal(logMsg)
	//	_, err := file.Write(strObj)
	//	fmt.Println("写了一条内容")
	//	if err != nil {
	//		fmt.Println("write file err", err)
	//	}
	//}
}
