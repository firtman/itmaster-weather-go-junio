package model

import "fmt"

type City struct {
	Id      int
	Name    string
	Country string
}

func (City) printCity() {
	fmt.Println("I'm a City")
}
