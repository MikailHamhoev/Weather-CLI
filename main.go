// weather-cli/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"weather-cli/api"
	"weather-cli/config"
	"weather-cli/display"
)

func main() {
	// Parse command line flags
	city := flag.String("city", "", "City name (required)")
	country := flag.String("country", "US", "Country code (default: US)")
	units := flag.String("units", "metric", "Units: metric or imperial")
	forecast := flag.Bool("forecast", false, "Show 5-day forecast")
	jsonOutput := flag.Bool("json", false, "Output as JSON")
	flag.Parse()

	// Validate required flags
	if *city == "" {
		fmt.Println("Error: City is required")
		fmt.Println("Usage: weather -city=<city> [-country=<code>] [-units=metric|imperial] [-forecast] [-json]")
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Create weather client
	client := api.NewWeatherClient(cfg.APIKey)

	// Fetch weather data
	if *forecast {
		forecastData, err := client.GetForecast(*city, *country, *units)
		if err != nil {
			log.Fatal("Failed to get forecast:", err)
		}

		if *jsonOutput {
			display.PrintJSON(forecastData)
		} else {
			display.PrintForecast(forecastData, *units)
		}
	} else {
		weather, err := client.GetCurrentWeather(*city, *country, *units)
		if err != nil {
			log.Fatal("Failed to get weather:", err)
		}

		if *jsonOutput {
			display.PrintJSON(weather)
		} else {
			display.PrintCurrentWeather(weather, *units)
		}
	}
}
