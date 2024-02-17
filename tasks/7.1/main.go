package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// синхронизация с использованием sync.Mutex
func main() {
	m := make(map[int]int)
	var mutex sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(m, &mutex, &wg)
	}

	wg.Wait()
	fmt.Println(m)
}

func write(m map[int]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	key := rand.Intn(100)
	value := rand.Intn(100) + 100
	mutex.Lock()
	defer mutex.Unlock()
	m[key] = value
}
