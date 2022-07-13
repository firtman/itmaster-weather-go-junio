package api

import (
	"fmt"
	"net/http"

	"andreani.com/go/weathergo/model"
)

const (
	apiURL = "https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v"
	apiKey = "85ad76cb241ddd400e9c40d1b59c5f74"
)

func GetWeather(cityName string) (*model.WeatherCity, error) {
	url := fmt.Sprintf(apiURL, cityName, apiKey)

	response, err := http.Get(url)

	if err != nil {
		// We have some response available
		wc := model.WeatherCity{}

		fmt.Println(response)

		return &wc, nil
	} else {
		return nil, err
	}
}
