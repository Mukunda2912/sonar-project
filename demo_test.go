package main

import (
    "testing"
)

func TestCategorizeTemperature(t *testing.T) {
    // Test cases with different temperatures and expected categories
    testCases := []struct {
        temperature float64
        expected    string
    }{
        {25.5, "Mild"},
        {-15.0, "Very Cold"},
        {8.7, "Cold"},
        {60.2, "Extremely Hot"},
    }

    // Iterate through the test cases and run the tests
    for _, tc := range testCases {
        result := categorizeTemperature(tc.temperature)
        if result != tc.expected {
            t.Errorf("Temperature %.1f°C: expected %s, got %s", tc.temperature, tc.expected, result)
        }
    }
}
