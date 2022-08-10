package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"andreani.com/go/weathergo/api"
)

func main() { // main goroutine
	fmt.Println("Weather Go 1.0 - Welcome")

	citiesPtr := flag.String("cities", "", "List of comma-separated cities")
	inputPtr := flag.String("input", "std", "How to get the cities to process, valid values: std, scan")
	outputPtr := flag.String("output", "std", "How to output the weather data, valid values: std, json")
	flag.Parse()

	var outputFunction func(city string)

	switch *outputPtr {
	case "std":
		outputFunction = processCity
	case "json":
		outputFunction = processCityToJson
	}

	switch *inputPtr {
	case "std":
		cities := strings.Split(*citiesPtr, ",")
		var wg sync.WaitGroup

		for _, city := range cities { // for each
			wg.Add(1)
			fmt.Printf("Looking for weather in %v ⏳\n", city) // sync
			go func(currentCity string) {                     // goroutine en un bloque anónimo
				outputFunction(currentCity)
				wg.Done() // wait synchronously to the whole group
			}(city)
		}
		wg.Wait()
	case "scan":
		fmt.Print("Enter the city name: ")
		var userCity string
		fmt.Scan(&userCity)
		outputFunction(userCity)
	}

	// time.Sleep(2 * time.Second) // espero en el main goroutine 3 segundos
}

func processCityToJson(city string) {
	wc, err := api.GetWeather(city) // sincrónico
	if err != nil {
		fmt.Println("We couldn't fetch the weather")
	} else {
		bytes, err := json.Marshal(wc)
		if err != nil {
			fmt.Println("Invalid JSON generation")
		} else {
			if _, err := os.Stat("output"); os.IsNotExist(err) {
				os.Mkdir("output", 0)
			}
			filename := fmt.Sprintf("output/%v.json", city)
			os.WriteFile(filename, bytes, 0)
		}
	}
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
