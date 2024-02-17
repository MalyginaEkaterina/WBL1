package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	//получаем значение для N
	dur := flag.Int("d", 1, "input duration in seconds")
	// делаем разбор командной строки
	flag.Parse()

	ch := make(chan int64)
	//добавляем в WaitGroup, чтобы дождаться завершения работы горутины
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		//цикл будет работать, пока не закончатся данные в закрытом канале
		for data := range ch {
			fmt.Println(data)
		}
		wg.Done()
	}()

	write(*dur, ch)
	//дожидаемся завершения работы горутины
	wg.Wait()
}

func write(dur int, out chan int64) {
	//создаем таймер, чтобы завершить отправку данных по истечении заданного времени
	timer := time.NewTimer(time.Duration(dur) * time.Second)
	var i int64
	for {
		select {
		case <-timer.C:
			close(out)
			return
		default:
			out <- i
			i++
		}
	}
}
