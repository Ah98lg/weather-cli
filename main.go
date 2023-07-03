package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type WeatherResponse struct {
	Location struct{
		Name string `json:"name"`
		Region string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Forecast struct{
		ForecastDay []struct{
			Hour []struct{
				TimeEpoch int64 `json:"time_epoch"`
				Temp_C float64 `json:"temp_c"`
				Condition struct{
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main(){

	q := "Natal"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	response, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=f7ac92d25ce54a288ec13333230307&q=" + q + "&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("Weather API server is not responding")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var weather WeatherResponse

	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, hours := weather.Location, weather.Forecast.ForecastDay[0].Hour


	fmt.Printf("%s, %s - %s\n", location.Name, location.Region, location.Country)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()){
			continue
		}



		message := fmt.Sprintf(
			"%s - %.0f°C, %.0f%%, %s \n",
			date.Format("15:04"),
			hour.Temp_C,
			hour.ChanceOfRain,
			hour.Condition.Text,
	)

		if hour.ChanceOfRain >= 50 {
			color.Red(message)
		} else {
			fmt.Print(message)
		}
	
	}

}