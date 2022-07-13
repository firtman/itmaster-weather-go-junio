package main

import (
	"fmt"

	"andreani.com/go/weathergo/api"
)

func main() {
	fmt.Println("Weather Go 1.0 - Welcome")
	wc, err := api.GetWeather("Madrid")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The temperature for %v is %.1f⁰C\n",
			wc.Name, wc.Temperature.ToCelsius())
	}
}
