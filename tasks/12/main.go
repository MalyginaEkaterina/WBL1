package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(getSet(strings))
}

func getSet(strings []string) map[string]struct{} {
	res := make(map[string]struct{})
	for i := 0; i < len(strings); i++ {
		res[strings[i]] = struct{}{}
	}
	return res
}
