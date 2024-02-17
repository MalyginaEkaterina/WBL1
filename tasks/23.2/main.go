package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

// Удаление без сохранения порядка
func main() {
	arr := []int{3, 5, 6, 7, 8, 9}
	arr, _ = deleteElement(arr, 0)
	arr2 := []int{3, 5, 6, 7, 8, 9}
	arr2, _ = deleteElement(arr2, 2)
	arr3 := []int{3, 5, 6, 7, 8, 9}
	arr3, _ = deleteElement(arr3, 5)
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// Ставим на место i-го последний элемент, убираем последний
func deleteElement(arr []int, i int) ([]int, error) {
	if i < 0 || i > len(arr)-1 {
		return arr, fmt.Errorf("wrong value for i")
	}
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1], nil
}
