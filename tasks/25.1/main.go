package main

import (
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	sleep(10)
}

func sleep(sec int) {
	start := time.Now()
	finish := start.Add(time.Duration(sec) * time.Second)
	for {
		//проверяем, прошло ли заданное количество секунд
		if finish.Before(time.Now()) {
			return
		}
	}
}
