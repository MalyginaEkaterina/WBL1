package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

func main() {
	a := Action{
		Human: Human{name: "John"},
		act:   "read",
	}
	a.do()
}

type Human struct {
	name string
}

func (h Human) getName() string {
	return fmt.Sprintf("name = %s", h.name)
}

func (h Human) getAction() string {
	return "nothing"
}

type Action struct {
	//Встраивание типа в структуру
	Human
	act string
}

func (a Action) getAction() string {
	return a.act
}

func (a Action) do() {
	fmt.Printf("do %s for %s", a.getAction(), a.getName())
}
