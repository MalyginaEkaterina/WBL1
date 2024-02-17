package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// множества представлены мапами
func main() {
	m1 := map[int]struct{}{3: {}, 4: {}, 5: {}, 9: {}, 13: {}, 14: {}, 15: {}, 16: {}}
	m2 := map[int]struct{}{1: {}, 5: {}, 16: {}, 9: {}, 13: {}, 17: {}}
	fmt.Println(getIntersection(m1, m2))
	fmt.Println(getIntersection(m2, m1))
}

func getIntersection(m1, m2 map[int]struct{}) map[int]struct{} {
	//проходим по мапе с меньшей длиной, чтобы сложность была меньше
	res := make(map[int]struct{})
	if len(m2) < len(m1) {
		m1, m2 = m2, m1
	}
	for k := range m1 {
		if _, ok := m2[k]; ok {
			res[k] = struct{}{}
		}
	}
	return res
}
