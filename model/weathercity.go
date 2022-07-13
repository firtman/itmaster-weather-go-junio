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

func (wc WeatherCity) String() string {
	return fmt.Sprintf("WEATHERCITY!!! %v: \nCity: %v\nWeather: %v",
		wc.timestamp, wc.City, wc.Weather)
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

	fmt.Println(t)
}
