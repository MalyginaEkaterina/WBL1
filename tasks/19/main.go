package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	fmt.Println(turnString("главрыба"))
}

func turnString(s string) string {
	//преобразуем строку в байт рун, чтобы работать с Unicode-символами
	runes := []rune(s)
	l := 0
	r := len(runes) - 1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}
