package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

*/

func main() {
	numWorkers := flag.Int("w", 1, "input number of workers")
	// делаем разбор командной строки
	flag.Parse()

	//канал для записи, ненулевое капасити выбран для тестирования корректного завершения работы
	out := make(chan interface{}, 100)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go func() {
			//цикл будет работать, пока не закончатся данные в закрытом канале
			for data := range out {
				//синхронизируем вывод в stdout, т.к. нет гарантий на запись из разных горутин одновременно
				mutex.Lock()
				fmt.Println(data)
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	//Обрабатываем сигналы для корректного завершения
	signalCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	//после сигнала канал out будет закрыт и программа будет ожидать,
	//пока воркеры обработают оставшиеся данные в канале, после чего завершится
	writer(signalCtx, out)
	wg.Wait()
}

func writer(ctx context.Context, out chan interface{}) {
	for {
		select {
		case <-ctx.Done():
			//добавляем последнюю фразу для проверки вывода всех данных в канале
			out <- "the end"
			close(out)
			return
		default:
			out <- generateRandString()
		}
	}
}

func generateRandString() string {
	randLen := rand.Intn(100) + 1

	arr := make([]byte, randLen)
	for i := 0; i < len(arr); i++ {
		arr[i] = byte(rand.Intn('z'-'a'+1) + 'a')
	}
	return string(arr)
}
