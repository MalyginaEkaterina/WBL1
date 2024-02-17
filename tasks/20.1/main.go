package main

import "fmt"

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// Решение для строк только из символов ASCII
// Сначала переворачиваем всю строку, затем обратно переворачиваем символы в словах
func main() {
	fmt.Println(turnWords("snow dog sun"))
}

func turnWords(s string) string {
	sb := []byte(s)
	reverse(sb, 0, len(sb)-1)

	//Теперь слова в правильном порядке, далее идем по словам и переворачиваем обратно символы в каждом слове
	l := 0
	r := 0
	for r < len(sb) {
		if sb[r] == ' ' {
			if r > l {
				reverse(sb, l, r-1)
			}
			r++
			l = r
		} else {
			r++
		}
	}
	if r > l {
		reverse(sb, l, r-1)
	}
	return string(sb)
}

func reverse(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
