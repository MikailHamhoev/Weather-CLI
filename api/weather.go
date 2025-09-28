// weather-cli/api/weather.go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type WeatherClient struct {
	apiKey string
	client *http.Client
}

type CurrentWeather struct {
	Location    string    `json:"location"`
	Country     string    `json:"country"`
	Temperature float64   `json:"temperature"`
	FeelsLike   float64   `json:"feels_like"`
	Humidity    int       `json:"humidity"`
	Pressure    int       `json:"pressure"`
	WindSpeed   float64   `json:"wind_speed"`
	WindDeg     int       `json:"wind_deg"`
	Weather     string    `json:"weather"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Sunrise     time.Time `json:"sunrise"`
	Sunset      time.Time `json:"sunset"`
	Time        time.Time `json:"time"`
	Units       string    `json:"units"`
}

type ForecastDay struct {
	Date        string  `json:"date"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Weather     string  `json:"weather"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Icon        string  `json:"icon"`
}

type ForecastResponse struct {
	Location string        `json:"location"`
	Country  string        `json:"country"`
	Days     []ForecastDay `json:"days"`
	Units    string        `json:"units"`
}

type OpenWeatherResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Dt int64 `json:"dt"`
}

type OpenWeatherForecastResponse struct {
	List []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"list"`
	City struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"city"`
}

func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (w *WeatherClient) GetCurrentWeather(city, country, units string) (*CurrentWeather, error) {
	// Build URL
	query := url.Values{}
	query.Add("q", fmt.Sprintf("%s,%s", city, country))
	query.Add("appid", w.apiKey)
	query.Add("units", units)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?%s", query.Encode())

	// Make request
	resp, err := w.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	// Parse response
	var data OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Convert to our format
	if len(data.Weather) == 0 {
		return nil, fmt.Errorf("no weather data in response")
	}

	weather := &CurrentWeather{
		Location:    data.Name,
		Country:     data.Sys.Country,
		Temperature: data.Main.Temp,
		FeelsLike:   data.Main.FeelsLike,
		Humidity:    data.Main.Humidity,
		Pressure:    data.Main.Pressure,
		WindSpeed:   data.Wind.Speed,
		WindDeg:     data.Wind.Deg,
		Weather:     data.Weather[0].Main,
		Description: data.Weather[0].Description,
		Icon:        data.Weather[0].Icon,
		Sunrise:     time.Unix(data.Sys.Sunrise, 0),
		Sunset:      time.Unix(data.Sys.Sunset, 0),
		Time:        time.Unix(data.Dt, 0),
		Units:       units,
	}

	return weather, nil
}

func (w *WeatherClient) GetForecast(city, country, units string) (*ForecastResponse, error) {
	// Build URL
	query := url.Values{}
	query.Add("q", fmt.Sprintf("%s,%s", city, country))
	query.Add("appid", w.apiKey)
	query.Add("units", units)
	query.Add("cnt", "40") // 5 days * 8 readings per day = 40

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?%s", query.Encode())

	// Make request
	resp, err := w.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	// Parse response
	var data OpenWeatherForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Group by day
	daysByDate := make(map[string]*ForecastDay)

	for _, item := range data.List {
		date := time.Unix(item.Dt, 0).Format("2006-01-02")

		if day, exists := daysByDate[date]; exists {
			// Update min/max temps
			if item.Main.TempMin < day.TempMin {
				day.TempMin = item.Main.TempMin
			}
			if item.Main.TempMax > day.TempMax {
				day.TempMax = item.Main.TempMax
			}
		} else {
			if len(item.Weather) == 0 {
				continue
			}

			daysByDate[date] = &ForecastDay{
				Date:        date,
				TempMin:     item.Main.TempMin,
				TempMax:     item.Main.TempMax,
				Weather:     item.Weather[0].Main,
				Description: item.Weather[0].Description,
				Humidity:    item.Main.Humidity,
				WindSpeed:   item.Wind.Speed,
				Icon:        item.Weather[0].Icon,
			}
		}
	}

	// Convert map to slice (max 5 days)
	var days []ForecastDay
	for _, day := range daysByDate {
		days = append(days, *day)
		if len(days) >= 5 {
			break
		}
	}

	forecast := &ForecastResponse{
		Location: data.City.Name,
		Country:  data.City.Country,
		Days:     days,
		Units:    units,
	}

	return forecast, nil
}
