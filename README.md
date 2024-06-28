## Weather Forecast Application

This Go application fetches and displays the current and hourly weather conditions for a specified location using the WeatherAPI. It highlights significant weather changes with colored output for better readability.

### Features

- Fetches current weather data and hourly forecast.
- Displays temperature, weather condition, and chance of rain.
- Highlights significant weather changes using colored output.
- Handles time zone formatting for accurate local times.

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/weather-forecast-app.git
    cd weather-forecast-app
    ```

2. Install dependencies:
    ```sh
    go get github.com/fatih/color
    ```

3. Set your WeatherAPI key in the code:
    ```go
    // Replace 'your_api_key_here' with your actual WeatherAPI key
    res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=your_api_key_here&q=" + q + "&days=1")
    ```

### Usage

Run the application with a specified location:
```sh
go run main.go "Location"
```
