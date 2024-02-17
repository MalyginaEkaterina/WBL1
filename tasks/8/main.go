package main

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	a := int64(340)
	fmt.Printf("%064b\n", a)
	a2, _ := changeBit(a, 6, 0)
	fmt.Printf("%064b\n", a2)
	a3, _ := changeBit(a, 0, 1)
	fmt.Printf("%064b\n", a3)

	b := int64(-340)
	fmt.Printf("%064b\n", uint64(b))
	b2, _ := changeBit(b, 62, 0)
	fmt.Printf("%064b\n", uint64(b2))
	b3, _ := changeBit(b, 0, 1)
	fmt.Printf("%064b\n", uint64(b3))
}

func changeBit(a int64, i int, bit int) (int64, error) {
	//так как у нас int64, а не uint64, то решено не разрешать менять 63й бит, т.к. он отвечает за знак
	if i < 0 || i > 62 {
		return 0, fmt.Errorf("i must be from 0 to 62")
	}
	if bit == 1 {
		// (1 << i) - число с 1 только в i-ом бите
		return a | (1 << i), nil
	} else if bit == 0 {
		// ^(1 << i) - инвертируем биты числа с 1 только в i-ом бите
		return a & (^(1 << i)), nil
	} else {
		return 0, fmt.Errorf("bit must be 0 or 1")
	}
}
