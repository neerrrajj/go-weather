package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// TO DO : add feels like to the output

type Weather struct {
	Location struct {
		Name           string `json:"name"`
		Country        string `json:"country"`
		TimeZone       string `json:"tz_id"`
		LocalTimeEpoch int64  `json:"localtime_epoch"`
		LocalTime      string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				Time      string  `json:"time"`
				Tempc     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Bangalore"
	if len(os.Args) >= 2 {
		q = url.QueryEscape(strings.Join(os.Args[1:], " "))
	}
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=b657cb636ca34cd3b3c83301242706&q=" + q + "&days=1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	city, country, temp, condition := weather.Location.Name, weather.Location.Country, weather.Current.TempC, weather.Current.Condition.Text
	hours := weather.Forecast.ForecastDay[0].Hour
	line := strings.Repeat("-", max(36, len(city+country+fmt.Sprintf("%.0f", temp)+condition)+8))

	fmt.Println(line)
	fmt.Printf("%s, %s: %.0fC, %s\n", city, country, temp, condition)
	// fmt.Printf("Time: %v\n", time.Unix(weather.Location.LocalTimeEpoch, 0))
	loc, err := time.LoadLocation(weather.Location.TimeZone)
	if err != nil {
		panic(err)
	}
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04", weather.Location.LocalTime, loc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Time: %v\n", parsedTime.Format("02-01-2006 15:04:05 -0700 MST"))
	fmt.Println(line)

	for _, hour := range hours {
		// date := time.Unix(hour.TimeEpoch, 0)
		hourmin, err := time.ParseInLocation("2006-01-02 15:04", hour.Time, loc)
		if err != nil {
			panic(err)
		}

		if hourmin.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %.0f%% rain, %s\n", hourmin.Format("15:04"), hour.Tempc, hour.ChanceOfRain, hour.Condition.Text)

		if hour.Tempc > 30 && hour.ChanceOfRain > 50 {
			color.Yellow((message))
		} else if hour.ChanceOfRain > 50 {
			color.Blue(message)
		} else if hour.Tempc > 30 {
			color.Red(message)
		} else {
			fmt.Print(message)
		}
	}
}
