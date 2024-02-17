package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	//Обрабатываем сигналы для завершения
	signalCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(5)
	go work1(signalCtx, &wg)
	stopCh := make(chan struct{})
	go work2(stopCh, &wg)
	go work3(&wg)
	dataCh := make(chan int)
	go work4(dataCh, &wg)
	var isStop atomic.Bool
	go work5(&isStop, &wg)

	go sendStop(stopCh)
	go sendData(dataCh)
	go changeIsStop(&isStop)

	wg.Wait()
}

// Завершение работы горутины с отменой контекста
func work1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("Stop work1")
			return
		}
	}
}

// Завершение работы горутины с закрытием или получением сигнала в канал
func work2(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			log.Println("Stop work2")
			return
		}
	}
}

// Завершение работы горутины через заданное время по таймеру
func work3(wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer.C:
			log.Println("Stop work3")
			return
		}
	}
}

// Завершение работы горутины после вычитывания данных из канала и его закрытия
func work4(dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		log.Println(data)
	}
	log.Println("Stop work4")
}

// Завершение работы горутины при изменении потокобезопасного флага
func work5(isStop *atomic.Bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if isStop.Load() {
			log.Println("Stop work5")
			return
		}
	}
}

func sendStop(stopCh chan struct{}) {
	time.Sleep(5 * time.Second)
	close(stopCh)
}

func sendData(dataCh chan int) {
	for i := 0; i < 7; i++ {
		dataCh <- i
	}
	close(dataCh)
}

func changeIsStop(isStop *atomic.Bool) {
	time.Sleep(3 * time.Second)
	isStop.Store(true)
}
