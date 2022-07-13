package model

import "fmt"

type City struct {
	Id      int
	Name    string
	Country string
}

func (city City) String() string {
	return fmt.Sprintf("City %v", city.Name)
}
