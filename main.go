package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var scan = bufio.NewScanner(os.Stdin)

func main() {
	for {
		color.White("----- Temperature Conveter -----")
		makeDecision(takeChoiceFromUser())
	}
}

func makeDecision(choice string) {
	switch choice {
	case "1":
		calculateFromCelsius()
	case "2":
		calculateFromFahrenheit()
	case "3":
		calculateFromKelvin()
	case "4":
		color.Green("\nThank you for using my application...")
		os.Exit(0)
	default:
		color.Red("\nPlease choose available options from menu\n\n")
	}
}

func calculateFromCelsius() {
	c := takeInputFromUser("Celsius")

	f := getFahrenheitFromCelsius(c)
	k := c + 273.15

	color.Yellow("\nInput Celsius value is %.2f\n", c)
	color.Yellow("Fahrenheit is %.2f and Kelvin is %.2f\n\n", f, k)
}

func calculateFromFahrenheit() {
	f := takeInputFromUser("Fahrenheit")

	c := (f - 32) * 5.0 / 9.0
	k := c + 273.15

	color.Yellow("Input Fahrenheit value is %.2f\n", f)
	color.Yellow("Celsius is %.2f and Kelvin is %.2f\n\n", c, k)
}

func calculateFromKelvin() {
	k := takeInputFromUser("Kelvin")

	c := k - 273.15
	// first converter kelvin to celsius and celsius to fahrenheit
	f := getFahrenheitFromCelsius(c)

	color.Yellow("Input Kevlin value is %.2f\n", k)
	color.Yellow("Celsius is %.2f and Fahrenheit is %.2f\n\n", c, f)
}

func takeChoiceFromUser() string {
	color.Cyan("\n1. Enter temperature in Celsius")
	color.Cyan("2. Enter temperature in Fahrenheit")
	color.Cyan("3. Enter temperature in Kelvin")
	color.Cyan("4. Exit from app")
	fmt.Println("\nEnter your choice")
	scan.Scan()
	return scan.Text()
}

func takeInputFromUser(input string) float64 {
	fmt.Println("\nEnter", input)
	scan.Scan()
	val := scan.Text()
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		color.Red("Invalid input. Please enter a valid number.")
	}
	return i
}

func getFahrenheitFromCelsius(c float64) float64 {
	return (9.0/5.0)*c + 32
}
