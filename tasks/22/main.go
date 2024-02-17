package main

import (
	"fmt"
	"math/big"
	"os"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

// Так как нет ограничения сверху, то использована библиотека math/big
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Error, enter operands and operation")
		os.Exit(1)
	}
	var a big.Int
	a.SetString(os.Args[1], 10)
	var b big.Int
	b.SetString(os.Args[3], 10)
	var c big.Int
	switch os.Args[2] {
	case "+":
		c.Add(&a, &b)
	case "-":
		c.Sub(&a, &b)
	case "*":
		c.Mul(&a, &b)
	case "/":
		c.Div(&a, &b)
	}
	fmt.Println(c.Text(10))
}
