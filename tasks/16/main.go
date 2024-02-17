package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	arr := []int{5, 8, 9, 15, 8, 16, 12, 6, 10, 20, 3, 9, 11, 1}
	qsort(arr)
	fmt.Println(arr)
}

func qsort(arr []int) {
	if len(arr) == 0 {
		return
	}
	iter(arr, 0, len(arr)-1)
}

func iter(arr []int, l, r int) {
	//обрабатываем отдельно массивы длины 1 и 2
	if l == r {
		return
	}
	if r-l == 1 {
		if arr[l] > arr[r] {
			arr[l], arr[r] = arr[r], arr[l]
		}
		return
	}
	p := partition(arr, l, r)
	iter(arr, l, p)
	iter(arr, p+1, r)
}

// выбираем элемент в середине массива, меняем элементы так, чтобы слева были
// меньше или равные ему, справа больше либо равные ему и возвращаем индекс(такой, что слева <= выбранному значению, справа >=)
func partition(arr []int, i, j int) int {
	pivot := arr[(i+j)/2]
	l := i
	r := j
	for l <= r {
		if arr[l] < pivot {
			l++
		} else if arr[r] > pivot {
			r--
		} else {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return r
}
