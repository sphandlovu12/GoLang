package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a mathematical expression (e.g., 5 + 3): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Calculator exiting. Goodbye!")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid input. Please enter an expression like '5 + 3'.")
			continue
		}

		num1, err1 := strconv.ParseFloat(parts[0], 64)
		operator := parts[1]
		num2, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input. Please enter valid numbers.")
			continue
		}

		var result float64
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				continue
			}
			result = num1 / num2
		default:
			fmt.Printf("Invalid operator: %s. Valid operators are +, -, *, and /.\n", operator)
			continue
		}

		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}
}
