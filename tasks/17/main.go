package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	arr := []int{1, 3, 5, 6, 8, 11, 13, 15, 17, 48, 59, 100, 114, 150}
	fmt.Println(binarySearch(arr, 5))
}

// возврщает индекс искомого элемента, либо -1, если элемент отсутствует
func binarySearch(arr []int, n int) int {
	if len(arr) == 0 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	for r-l > 1 {
		midInd := l + (r-l)/2
		if arr[midInd] == n {
			return midInd
		} else if arr[midInd] > n {
			r = midInd
		} else {
			l = midInd
		}
	}
	if arr[l] == n {
		return l
	}
	if arr[r] == n {
		return r
	}
	return -1
}
