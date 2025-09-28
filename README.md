# Weather CLI

A command-line weather tool that fetches real-time data and forecasts from OpenWeatherMap.

## Setup

### 1. Get API Key
- Sign up at https://home.openweathermap.org/users/sign_up  
- Free tier: 60 calls/minute, 1M/month  

### 2. Configure API Key
**Option A (recommended):**
```bash
export WEATHER_API_KEY="your_api_key_here"
```

**Option B:**
```bash
mkdir -p ~/.weather-cli
echo "your_api_key_here" > ~/.weather-cli/config
chmod 600 ~/.weather-cli/config
```

## Usage

### Build
```bash
go mod tidy
go build -o weather
```

### Commands
```bash
# Current weather
./weather -city="London" -country="GB"

# Imperial units
./weather -city="New York" -country="US" -units=imperial

# 5-day forecast
./weather -city="Tokyo" -country="JP" -forecast

# JSON output
./weather -city="Paris" -country="FR" -json

# Combined options
./weather -city="Toronto" -country="CA" -units=imperial -forecast -json
```

## Sample Output
```
==================================================
🌍 Weather for London, GB
📅 Mon, Jan 15 14:30
--------------------------------------------------
☁️ Clouds
   broken clouds

🌡️  Temperature: 12.5°C
   Feels like: 11.0°C
💧 Humidity: 65%
📊 Pressure: 1013 hPa
💨 Wind: 5.2 m/s
   Direction: 240°
🌅 Sunrise: 07:45
🌇 Sunset: 16:30
==================================================
```

## Tips
- **Global install**: `sudo cp weather /usr/local/bin/`  
- **Run without building**: `go run main.go -city="Tokyo" -country="JP"`  
- **Batch check**: loop through cities with a shell script  

## Features Demonstrated
- External API integration  
- JSON parsing  
- Config management (env vars + file)  
- CLI flag parsing  
- Error handling  
- Formatted output  
- Unit conversion (metric/imperial)

MIT License