package model

import (
	"fmt"
)

type WeatherCity struct {
	// Embedding
	City
	Weather
	timestamp int
}

func Test() {
	t := WeatherCity{}
	t.Humidity = 34
	t.Name = "Rosario"
	t.Country = "Argentina"
	t.Temperature = 24
	t.timestamp = 234234234
	t.Weather.Condition = 34
	t.City.Name = ""
	t.printCity()

	fmt.Println(t)
}
