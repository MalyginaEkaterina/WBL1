package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	//с использованием побитового XOR
	a := 34
	b := 58
	// a^0 = a, a^a = 0
	b = a ^ b
	// a = a^a^b = b
	a = a ^ b
	// b = a^a^b^a^b = a
	b = a ^ b
	fmt.Printf("a := %d, b := %d\n", a, b)

	//через двойное присваивание
	a, b = b, a
	fmt.Printf("a := %d, b := %d\n", a, b)
}
