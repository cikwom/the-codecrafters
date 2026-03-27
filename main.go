package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Happy's Calculator")
	fmt.Println("Type 'help' to see commands")

	for {
		var command string
		var a, b float64

		fmt.Println("choose a command: ")
		_, err := fmt.Scanln(&command, &a, &b)

		if command == "quit" {
			fmt.Println("Goodbye")
			break
		}

		if command == "help" {
			fmt.Println("add <a> <b>")
			fmt.Println("sub <a> <b>")
			fmt.Println("mul <a> <b>")
			fmt.Println("div <a> <b>")
			fmt.Println("quit")
			continue
		}

		if err != nil {
			fmt.Println("Error: Invalid Input, ")
			continue
		}

		var result float64

		switch command {
		case "add":
			result = a + b

		case "sub":
			result = a - b

		case "mul":
			result = a * b

		case "div":
			if b == 0 {
				fmt.Println("Error: Cannot divide by zero")
				continue
			}
			result = a / b

		default:
			fmt.Println("Invalid command: Type help")
			continue
		}
		fmt.Println("Result:", result)
	}
}
