package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	out := make(chan int)
	qOut := make(chan int)

	//горутина читает числа из первого канала, отправляет их удвоенными во второй канал
	//после закрытия первого канала и обработки всех чисел из него закрывает второй канал
	go func() {
		for n := range out {
			qOut <- 2 * n
		}
		close(qOut)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	//горутина читает числа из второго канала и выводит их в stdout
	go func() {
		for n := range qOut {
			fmt.Println(n)
		}
		wg.Done()
	}()

	for i := 0; i < len(arr); i++ {
		out <- i
	}
	close(out)
	//ожидаем завершения только второй горутины, т.к. она всегда завершается после первой
	wg.Wait()
}
