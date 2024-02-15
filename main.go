package main

import (
	"fmt"
	"os"
	"strconv"
)

const KminusC float64 = 273.0

func KelvinToCelsius(temperatureC float64) float64 {
	return temperatureC - KminusC
}

func CelsiusToKelvin(temperatureK float64) float64 {
	return temperatureK + KminusC
}

func main() {
	args := os.Args[1:]

	option := args[0]

	if option == "-h" || option == "--help" {
		fmt.Printf("Converts Celsius to Kelvin and vice-versa.\n" +
				"USAGE:\n"+
				"\t-h or --help\tdisplay this help\n"+
				"\t-tok [temp]\tconverts temperature from Celsius to Kelvin\n" +
				"\t-toc [temp]\tconverts temperature from Kelvin to Celsius\n")
		return
	}

	value, _ := strconv.ParseFloat(args[1], 64)
	
	if option == "-toc" {
		finalValue := KelvinToCelsius(value)
		fmt.Printf("%.1f Kelvin is %.1f Celsius\n", value, finalValue)
		return
	}

	if option == "-tok" {
		finalValue := CelsiusToKelvin(value)
		fmt.Printf("%.1f Celsius is %.1f Kelvin\n", value, finalValue)
		return
	}
	
}
