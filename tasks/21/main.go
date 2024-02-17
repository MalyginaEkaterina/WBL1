package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Адаптер — позволяет объектам с разными интерфейсами работать вместе
func main() {
	cd := ClientData{
		name: "Ivan",
		age:  33,
	}
	PrintAnything(&cd)

	c := ClientFromDifferentSystem{
		firstName: "Ivan",
		lastName:  "Ivanov",
	}
	a := Adapter{client: &c}
	PrintAnything(&a)
}

// Интерфейс явно реализуется ClientData, а для ClientFromDifferentSystem через адаптер
type DataPrinter interface {
	PrintData()
}

type ClientData struct {
	name string
	age  int
}

func (c *ClientData) PrintData() {
	fmt.Printf("Our client: %s, age: %d\n", c.name, c.age)
}

func PrintAnything(p DataPrinter) {
	p.PrintData()
}

type ClientFromDifferentSystem struct {
	firstName string
	lastName  string
}

func (c *ClientFromDifferentSystem) PrintName() {
	fmt.Printf("Client name: %s %s\n", c.firstName, c.lastName)
}

type Adapter struct {
	client *ClientFromDifferentSystem
}

func (a *Adapter) PrintData() {
	a.client.PrintName()
}
