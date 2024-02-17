package main

import (
	"context"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	sleep(7)
}

func sleep(sec int) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	<-ctx.Done()
}
