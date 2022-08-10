package main

import (
	"fmt"
	"time"

	"andreani.com/go/weathergo/api"
)

func main() { // main couroutine
	fmt.Println("Weather Go 1.0 - Welcome")
	cities := []string{"Madrid", "Buenos Aires", "Rosario", "Moscow"}

	for _, city := range cities { // for each
		fmt.Printf("Looking for weather in %v ⏳\n", city) // sync
		go processCity(city)                              // async
	}

	time.Sleep(2 * time.Second) // espero en el main goroutine 3 segundos
}

func processCity(city string) {
	wc, err := api.GetWeather(city) // sincrónico
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The temperature for %v is %.1f⁰C\n",
			wc.Name, wc.Temperature.ToCelsius())
	}
}
