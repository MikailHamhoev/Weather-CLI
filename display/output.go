// weather-cli/display/output.go
package display

import (
	"encoding/json"
	"fmt"
	"strings"
	"weather-cli/api"
	"weather-cli/utils"
)

func PrintCurrentWeather(weather *api.CurrentWeather, units string) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("🌍 Weather for %s, %s\n", weather.Location, weather.Country)
	fmt.Printf("📅 %s %s\n", utils.FormatDate(weather.Time), utils.FormatTime(weather.Time))
	fmt.Println(strings.Repeat("-", 50))

	icon := utils.GetWeatherIcon(weather.Icon)
	fmt.Printf("%s %s\n", icon, weather.Weather)
	fmt.Printf("   %s\n", weather.Description)
	fmt.Println()

	fmt.Printf("🌡️  Temperature: %s\n", utils.FormatTemperature(weather.Temperature, units))
	fmt.Printf("   Feels like: %s\n", utils.FormatTemperature(weather.FeelsLike, units))
	fmt.Printf("💧 Humidity: %s\n", utils.FormatHumidity(weather.Humidity))
	fmt.Printf("📊 Pressure: %s\n", utils.FormatPressure(weather.Pressure))
	fmt.Printf("💨 Wind: %s\n", utils.FormatWindSpeed(weather.WindSpeed, units))
	fmt.Printf("   Direction: %d°\n", weather.WindDeg)
	fmt.Println()

	fmt.Printf("🌅 Sunrise: %s\n", utils.FormatTime(weather.Sunrise))
	fmt.Printf("🌇 Sunset: %s\n", utils.FormatTime(weather.Sunset))
	fmt.Println(strings.Repeat("=", 50))
}

func PrintForecast(forecast *api.ForecastResponse, units string) {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("📅 5-Day Forecast for %s, %s\n", forecast.Location, forecast.Country)
	fmt.Println(strings.Repeat("-", 60))

	for i, day := range forecast.Days {
		icon := utils.GetWeatherIcon(day.Icon)
		fmt.Printf("%d. %s %s\n", i+1, icon, day.Date)
		fmt.Printf("   🌡️  Min: %s, Max: %s\n",
			utils.FormatTemperature(day.TempMin, units),
			utils.FormatTemperature(day.TempMax, units))
		fmt.Printf("   ☁️  %s: %s\n", day.Weather, day.Description)
		fmt.Printf("   💧 Humidity: %s\n", utils.FormatHumidity(day.Humidity))
		fmt.Printf("   💨 Wind: %s\n", utils.FormatWindSpeed(day.WindSpeed, units))
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 60))
}

func PrintJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}
