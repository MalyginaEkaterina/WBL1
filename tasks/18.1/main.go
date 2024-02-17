package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Вариант с синхронизацией с помощью Mutex
func main() {
	c := counter{}

	var wg sync.WaitGroup
	for i := 4; i < 10; i++ {
		wg.Add(1)
		go increase(&c, i, &wg)
	}

	wg.Wait()
	fmt.Println(c.get())
}

func increase(c *counter, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < i; j++ {
		c.add()
	}
}

type counter struct {
	value int
	mutex sync.Mutex
}

func (c *counter) add() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// RWMutex работает чуть медленнее Mutex, а для очень коротких интервалов не имеет смысла
func (c *counter) get() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}
