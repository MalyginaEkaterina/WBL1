package main

import "fmt"

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// Выделяем массив байт длины, равной длине строки.
// Далее идем с конца строки и записываем слова в выделенный массив.
func main() {
	fmt.Println(turnWords("snow dog sun"))
}

func turnWords(s string) string {
	resB := make([]byte, len(s))
	writeInd := 0
	l := len(s) - 1
	r := len(s) - 1
	for l >= 0 {
		if s[l] == ' ' {
			//если перед пробелом было слово, записываем его
			if r > l {
				for j := l + 1; j <= r; j++ {
					resB[writeInd] = s[j]
					writeInd++
				}
			}
			//записываем пробел
			resB[writeInd] = s[l]
			writeInd++
			l--
			r = l
		} else {
			l--
		}
	}
	//записываем последнее слово
	if r > l {
		for j := 0; j <= r; j++ {
			resB[writeInd] = s[j]
			writeInd++
		}
	}
	return string(resB)
}
