package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	src := []int{2, 4, 6, 8, 10}
	//определяем количество горутин в зависимости от количества ядер
	n := runtime.NumCPU()
	if n > len(src) {
		n = len(src)
	}
	//определяем объем работы каждой горутины
	maxCount := len(src) / n
	if len(src)%n > 0 {
		maxCount++
	}
	res := make([]int, len(src))
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			l := j * maxCount
			r := (j+1)*maxCount - 1
			if r >= len(src) {
				r = len(src) - 1
			}
			for k := l; k <= r; k++ {
				//k в разных горутинах не пересекаются, поэтому синхронизация не нужна
				res[k] = src[k] * src[k]
			}
		}(i)
	}

	//ожидаем завершения работы всех горутин
	wg.Wait()

	for i := 0; i < len(src); i++ {
		fmt.Println(res[i])
	}
}
