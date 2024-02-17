package main

import "time"

/*
Реализовать собственную функцию sleep.
*/

func main() {
	sleep(5)
}

func sleep(sec int) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
}
