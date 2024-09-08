package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Sheep struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}
func (s Sheep) Speak() string {
	return "Beee"
}

func makeItSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	myDog := Dog{Name: "Rex"}
	myCat := Cat{Name: "Kitty"}
	mySheep := Sheep{Name: "Shaun"}
	makeItSpeak(myDog)
	makeItSpeak(myCat)
	makeItSpeak(mySheep)
}
