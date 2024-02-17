package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// Разбиваем строку на массив слов, его переворачиваем и собираем обратно в строку
func main() {
	fmt.Println(turnWords("snow dog sun"))
}

func turnWords(s string) string {
	arr := strings.Split(s, " ")
	l := 0
	r := len(arr) - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return strings.Join(arr, " ")
}
