package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
}

type Flyable interface {
	Fly()
}

func (p Person) Fly() {
	fmt.Println("I'm a person and I can fly")
}

func foo(f Flyable) {
	p, ok := f.(*Person)
	if !ok {
		fmt.Print("I'm not a person")
		return
	}
	p.Name = "foo" + p.Name
	fmt.Println(p.Name)
}

func main() {
	p := Person{Name: "John"}
	foo(&p)
	fmt.Print(p.Name)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
