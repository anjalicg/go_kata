package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9

}

func main() {
	// First command line parameter is the value in fahrenheit
	f, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("Celsius of %fF is %f \n", f, convertToCelsius(f))
	fmt.Println("Enter more temperature valuues in Fahrenheit")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nFahrenheit>")
		text, _ := reader.ReadString('\n')
		n, _ := strconv.ParseFloat(strings.TrimSpace(text), 64)
		fmt.Println(n)
		fmt.Printf("Celsius of %fF is %f \n", n, convertToCelsius(n))
	}
}
