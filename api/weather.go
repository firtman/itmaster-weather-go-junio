package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"andreani.com/go/weathergo/model"
)

const (
	apiURL = "https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v"
	apiKey = "85ad76cb241ddd400e9c40d1b59c5f74"
)

func GetWeather(cityName string) (*model.WeatherCity, error) {
	url := fmt.Sprintf(apiURL, cityName, apiKey)

	response, err := http.Get(url) // sincrónico hasta que los headers estén disponibles

	if err != nil {
		// hubo algún error de red
		return nil, err
	} else {
		if response.StatusCode == http.StatusOK {
			bytes, err := io.ReadAll(response.Body) // sincrónico hasta que estén todos los bytes
			if err != nil {
				return nil, err
			} else {
				// We have some response available
				wc := model.WeatherCity{}
				parseWeatherJSONPretty(bytes, &wc)
				return &wc, nil
			}
		} else {
			// return nil, errors.New(fmt.Sprintf("HTTP Status code: %v", response.StatusCode))
			// return nil, fmt.Errorf("HTTP Status code: %v", response.StatusCode)
			return nil, ApiError{code: response.StatusCode, message: "Error HTTP",
				extra: cityName}
		}
	}
}

func parseWeatherJSONPretty(bytes []byte, wc *model.WeatherCity) {
	var result OpenWeatherMapResponse
	json.Unmarshal(bytes, &result)
	wc.Name = result.Name
	wc.Id = result.Id
	wc.Temperature = model.DegreeK(result.Main.Temperature)
	wc.Humidity = int(result.Main.Humidity)
}

func parseWeatherJSON(bytes []byte, wc *model.WeatherCity) {
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)

	wc.Name = result["name"].(string)
	wc.Temperature = model.DegreeK(result["main"].(map[string]interface{})["temp"].(float64))

}
