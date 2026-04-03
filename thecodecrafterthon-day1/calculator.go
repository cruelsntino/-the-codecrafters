package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subtract(a, b float64) float64 {
	return float64(a - b)
}

func multiply(a, b float64) float64 {
	return float64(a) * float64(b)
}

func add(a, b float64) float64 {
	return float64(a) + float64(b)
}

func divide(a, b float64) float64 {
	return float64(a) / float64(b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Abascus Calculator")
	fmt.Println("Welcome: Type 'exit' to quit or 'help' to start")
	for {
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		if input == "exit" || input == "quit" {
			fmt.Println("Thanks for banking with us!")
			return

		}
		if input == "help" {
			fmt.Println("add <a> <b>   → addition")
			fmt.Println("sub <a> <b>   → subtraction")
			fmt.Println("mul <a> <b>   → multiplication")
			fmt.Println("div <a> <b>   → division")
			fmt.Println("'exit' to quit")
			continue
		}
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		command := strings.ToLower(parts[0])

		if len(parts) != 3 {
			fmt.Printf("Usage: %s <a> <b>\n", command)
			continue
		}

		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Not a digit")
			continue
		}

		if command != "help" && command != "add" && command != "sub" && command != "mul" && command != "div" && command != "exit" {
			fmt.Println("Unknown commands: type 'help' to continue")
			continue
		}

		switch command {
		case "add":
			fmt.Printf("%g + %g = %g\n", a, b, add(a, b))
		case "sub":
			fmt.Printf("%g - %g = %g\n", a, b, subtract(a, b))
		case "mul":
			fmt.Printf("%g * %g = %g\n", a, b, multiply(a, b))
		case "div":
			if b == 0 {
				fmt.Printf("%g is not divisible by %g\n", a, b)
				continue
			}
			fmt.Printf("%g / %g = %g\n", a, b, divide(a, b))
		default:
			fmt.Println("Invalid command: Type 'help' for guidelines")
		}
	}
}
