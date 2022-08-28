package bmi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func printWelcome() {
	fmt.Println("BMI Calculator")
	fmt.Println("--------------")
}

func getUserInput(promptText string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ := strconv.ParseFloat(userInput, 64)

	return value
}

func getUserMetrics() (float64, float64) {
	weight := getUserInput("Please enter your weight (kg): ")
	height := getUserInput("Please enter your height (m): ")

	return weight, height
}

func calcBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func BMISample() {

	printWelcome()

	weight, height := getUserMetrics()

	bmi := calcBMI(weight, height)

	fmt.Printf("Your BMI: %.2f", bmi)
}
