package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// синхронизация с использованием канала
func main() {
	m := make(map[int]int)
	writeCh := make(chan pair)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(writeCh, &wg)
	}

	//запись данных в map происходит из одной горутины
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for p := range writeCh {
			m[p.key] = p.value
		}
	}()

	wg.Wait()
	close(writeCh)
	wg2.Wait()
	fmt.Println(m)
}

type pair struct {
	key   int
	value int
}

func write(writeCh chan pair, wg *sync.WaitGroup) {
	defer wg.Done()
	key := rand.Intn(100)
	value := rand.Intn(100) + 100
	writeCh <- pair{
		key:   key,
		value: value,
	}
}
