package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
}

func (h Human) Say(phrase string) {
	fmt.Println(phrase)
}

type Action struct {
	Human
	Action string
}

func main() {
	s := Action{
		Human: Human{
			Name:   "Ivan",
			Age:    20,
			Height: 190,
			Weight: 75,
		},
		Action: "Walk",
	}

	//Метод Say наследуется от Human
	s.Say("Hello World!")
}
