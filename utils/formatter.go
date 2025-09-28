// weather-cli/utils/formatter.go
package utils

import (
	"fmt"
	"time"
)

func FormatTemperature(temp float64, units string) string {
	symbol := "°C"
	if units == "imperial" {
		symbol = "°F"
	}
	return fmt.Sprintf("%.1f%s", temp, symbol)
}

func FormatWindSpeed(speed float64, units string) string {
	unit := "m/s"
	if units == "imperial" {
		unit = "mph"
	}
	return fmt.Sprintf("%.1f %s", speed, unit)
}

func FormatPressure(pressure int) string {
	return fmt.Sprintf("%d hPa", pressure)
}

func FormatHumidity(humidity int) string {
	return fmt.Sprintf("%d%%", humidity)
}

func FormatTime(t time.Time) string {
	return t.Format("15:04")
}

func FormatDate(t time.Time) string {
	return t.Format("Mon, Jan 2")
}

func GetWeatherIcon(icon string) string {
	// Map OpenWeather icons to emojis
	iconMap := map[string]string{
		"01d": "☀️", // clear sky day
		"01n": "🌙",  // clear sky night
		"02d": "⛅",  // few clouds day
		"02n": "☁️", // few clouds night
		"03d": "☁️", // scattered clouds
		"03n": "☁️", // scattered clouds
		"04d": "☁️", // broken clouds
		"04n": "☁️", // broken clouds
		"09d": "🌧️", // shower rain
		"09n": "🌧️", // shower rain
		"10d": "🌦️", // rain day
		"10n": "🌧️", // rain night
		"11d": "⛈️", // thunderstorm
		"11n": "⛈️", // thunderstorm
		"13d": "❄️", // snow
		"13n": "❄️", // snow
		"50d": "🌫️", // mist
		"50n": "🌫️", // mist
	}

	if emoji, ok := iconMap[icon]; ok {
		return emoji
	}
	return "☀️" // default
}
