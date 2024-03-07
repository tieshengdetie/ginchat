package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
}

func (a *Animal) Eat(str string) {
	fmt.Println(str)
}

func main() {
	a := Animal{}
	pByte := "nihao"
	param := make([]reflect.Value, 1)
	param[0] = reflect.ValueOf(pByte)
	//param = append(param, )

	reflect.ValueOf(&a).MethodByName("Eat").Call(param)
	fmt.Println(reflect.TypeOf(pByte))
	//reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})

}
