package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// множества представлены слайсами
// предполагаем, что повторяющихся элементов в слайсах нет, так как это множества
func main() {
	a1 := []int{4, 5, 7, 8, 9, 12, 14, 15, 18}
	a2 := []int{1, 3, 6, 8, 14, 20}
	fmt.Println(getIntersection(a1, a2))
	fmt.Println(getIntersection(a2, a1))
}

func getIntersection(a1, a2 []int) []int {
	var res []int
	m := make(map[int]struct{})
	if len(a2) < len(a1) {
		a1, a2 = a2, a1
	}
	//так как не удастся сделать сложность меньше, чем n+m, то формируем мапу по меньшему слайсу,
	//чтобы сэкономить на промежуточной памяти
	for i := 0; i < len(a1); i++ {
		m[a1[i]] = struct{}{}
	}
	for i := 0; i < len(a2); i++ {
		if _, ok := m[a2[i]]; ok {
			res = append(res, a2[i])
		}
	}
	return res
}
