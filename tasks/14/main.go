package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	a := 5
	fmt.Println(returnType(a))
	b := "hello"
	fmt.Println(returnType(b))
	c := true
	fmt.Println(returnType(c))
	d := make(chan int)
	fmt.Println(returnType(d))
}

func returnType(a interface{}) string {
	t := reflect.TypeOf(a)
	if t.Kind() == reflect.Chan {
		//t.String() вернет chan T, но в задании указано channel
		return "channel"
	}
	return t.Name()
}
