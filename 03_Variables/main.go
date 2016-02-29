package main

import "fmt"

func main() {
	fmt.Print("Enter temprature in Fahrenheit: ")
	var tempratureFahren float64
	fmt.Scanf("%f", &tempratureFahren)
	tempratureCelsius := ((tempratureFahren -32) * 5/9)
	fmt.Println("Output in Celsius: ", tempratureCelsius)

}
