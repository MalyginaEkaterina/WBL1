package main

import "math/rand"

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

//В памяти будет оставаться большая строка, т.к. на нее ссылается justString,
//хотя нам дальше потребуется только ее начало.
//Нужно скопировать эти данные, тогда далее неиспользуемые данные будут удалены GC.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	dst := make([]byte, 100)
	copy(dst, v[:100])
	justString = string(dst)
}

func createHugeString(i int) []byte {
	arr := make([]byte, i)
	for j := 0; j < i; j++ {
		arr[j] = byte(rand.Intn('z'-'a'+1) + 'a')
	}
	return arr
}

func main() {
	someFunc()
	println(justString)
}
