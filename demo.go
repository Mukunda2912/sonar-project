package main

import (
    "fmt"
)

func categorizeTemperature(temp float64) string {
    if temp < -10 {
        return "Extremely Cold"
    } else if temp >= -10 && temp < 0 {
        return "Very Cold"
    } else if temp >= 0 && temp < 10 {
        return "Cold"
    } else if temp >= 10 && temp < 20 {
        return "Cool"
    } else if temp >= 20 && temp < 30 {
        return "Mild"
    } else if temp >= 30 && temp < 40 {
        return "Warm"
    } else if temp >= 40 && temp < 50 {
        return "Hot"
    } else if temp >= 50 {
        return "Extremely Hot"
    } else {
        return "Unknown"
    }
}

func main() {
    temperature := 25.5
    category := categorizeTemperature(temperature)
    fmt.Printf("The temperature %.1f°C is considered %s.\n", temperature, category)

    temperature = -15.0
    category = categorizeTemperature(temperature)
    fmt.Printf("The temperature %.1f°C is considered %s.\n", temperature, category)

    temperature = 8.7
    category = categorizeTemperature(temperature)
    fmt.Printf("The temperature %.1f°C is considered %s.\n", temperature, category)

    temperature = 60.2
    category = categorizeTemperature(temperature)
    fmt.Printf("The temperature %.1f°C is considered %s.\n", temperature, category)
}
