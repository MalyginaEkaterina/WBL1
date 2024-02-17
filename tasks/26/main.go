package main

import (
	"fmt"
	"unicode"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	fmt.Println(allUnique("abcd"))
	fmt.Println(allUnique("abCdefA"))
}

// Приводим символы к нижнему регистру, ищем в мапе.
// Если такой символ уже встречался ранее, возвращаем false, иначе добавляем в мапу и проверяем далее
func allUnique(s string) bool {
	m := make(map[rune]struct{})
	for _, r := range s {
		lr := unicode.ToLower(r)
		if _, ok := m[lr]; ok {
			return false
		}
		m[lr] = struct{}{}
	}
	return true
}
