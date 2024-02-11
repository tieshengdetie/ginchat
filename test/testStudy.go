package main

import "fmt"

func main() {
	//// 对数组做切片
	//array := [3]int{1, 2, 3} // array是数组
	//slice3 := array[1:3]     // slice3是切片
	//fmt.Println("slice3 type:", reflect.TypeOf(slice3))
	//fmt.Println("slice3=", slice3) // slice3= [2 3]
	//a := []int{1, 2}
	//b := append(a, 3)
	//
	//c := append(b, 4)
	//d := append(b, 5)
	//
	//fmt.Println(a, b, c[3], d[3])
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	//s = append(s, 4)
	//s = append(s, 5)
	//s = append(s, 6)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

}
