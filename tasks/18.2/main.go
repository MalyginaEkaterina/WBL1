package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Вариант с использованием atomic
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
	value atomic.Int32
}

func (c *counter) add() {
	c.value.Add(1)
}

func (c *counter) get() int32 {
	return c.value.Load()
}
