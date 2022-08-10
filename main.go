package main

import (
	"fmt"
	"os"
	"sync"

	"andreani.com/go/weathergo/api"
)

func main() { // main goroutine
	fmt.Println("Weather Go 1.0 - Welcome")

	cities := os.Args[1:]

	var wg sync.WaitGroup

	for _, city := range cities { // for each
		wg.Add(1)
		fmt.Printf("Looking for weather in %v ⏳\n", city) // sync
		go func(currentCity string) {                     // goroutine en un bloque anónimo
			processCity(currentCity)
			wg.Done() // wait synchronously to the whole group
		}(city)
	}
	wg.Wait()
	// time.Sleep(2 * time.Second) // espero en el main goroutine 3 segundos
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
