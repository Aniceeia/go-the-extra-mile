package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput(str string, isDigit bool) interface{} {
	for {
		fmt.Println(str)
		var rawString string
		fmt.Scanln(&rawString)
		rawString = strings.TrimSpace(rawString)
		if isDigit {
			val, err := strconv.ParseFloat(rawString, 64)
			if err == nil {
				return val
			}
			fmt.Println("Invalid input")
		} else {
			correctOperations := []string{"+", "-", "*", "/"}
			for _, val := range correctOperations {
				if rawString == val {
					return rawString
				}
			}
			fmt.Println("Invalid input")
		}
	}
}

func main() {
	left := readInput("Input left operand:", true).(float64)
	operation := readInput("Input operation", false).(string)
	right := readInput("Input right operand:", true).(float64)
	var result float64
	switch operation {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right != 0 {
			result = left / right
		} else {
			fmt.Println("ERROR: division by zero")
		}
	}
	fmt.Printf("%.3f", result)
}
