package main

import (
	"fmt"
	"runtime"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+4^2+6^2….) с использованием конкурентных вычислений.
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
	if len(src)/n > 0 {
		maxCount++
	}

	//канал для отправки промежуточных результатов подсчета сумм
	sumCh := make(chan int)
	for i := 0; i < n; i++ {
		go func(j int) {
			l := j * maxCount
			r := (j+1)*maxCount - 1
			if r >= len(src) {
				r = len(src) - 1
			}
			//независимо считаем суммы каждого куска массива
			sum := 0
			for k := l; k <= r; k++ {
				sum += src[k] * src[k]
			}
			sumCh <- sum
		}(i)
	}

	//считаем финальную сумму
	sum := 0
	count := 0
	for add := range sumCh {
		sum += add
		count++
		if count == len(src) {
			break
		}
	}
	fmt.Println(sum)
}
