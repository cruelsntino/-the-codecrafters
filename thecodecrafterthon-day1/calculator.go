package main

import (
	"fmt"
	"strings"
)

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var a, b int
	var op string
	var cont string
	fmt.Println("Abascus Calculator")
	fmt.Println("Welcome: Type 'quit' to exit or 'help' to start")

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&op)

		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		if op == "quit" {
			fmt.Println("Thanks for banking with us")
			return
		}

		if op == "help" {
			fmt.Println("Choose Operation:")
			fmt.Println(" add <a> <b> → addition")
			fmt.Println(" sub <a> <b> → subtraction")
			fmt.Println(" mul <a> <b> → multiplication")
			fmt.Println(" div <a> <b> → division")
			fmt.Println(" quit → exit")
			continue
		}
		if op != "add" && op != "sub" && op != "mul" && op != "div" {
			fmt.Println("Invalid operation. Type 'help'")
			continue
		}

	begin:
		fmt.Printf("Choose Operation\n  add <a> <b>   → addition\n  sub <a> <b>   → subtraction\n  mul <a> <b>   → multiplication\n  div <a> <b>   → division\n  quit\n")

		fmt.Println("Enter First Input:")
		_, err = fmt.Scan(&a)
		if err != nil {
			fmt.Printf("Input a valid number\n")
			fmt.Scanln()
			goto begin
		}

		fmt.Println("Enter Second Input:")
		_, err = fmt.Scan(&b)
		if err != nil {
			fmt.Printf("Input a valid number\n")
			fmt.Scanln()
			goto begin
		}

		switch op {
		case "add":
			fmt.Printf("The addition of %d + %d: '%d'\n", a, b, add(a, b))

			// fmt.Println("Type ok to start another calculation")

		case "sub":
			fmt.Printf("The subtraction of %d - %d: '%d'\n", a, b, subtract(a, b))
			// fmt.Println("Come on")

		case "mul":
			fmt.Printf("The multiplication of %d * %d: %d\n", a, b, multiply(a, b))
			// fmt.Println("Thats IT; you can do better")

		case "div":
			if b == 0 {
				fmt.Println("Error: Indivisible by zero")

			}
			fmt.Printf("The Division of %d / %d: %.2f\n", a, b, divide(a, b))

			// fmt.Println("That Tickles")

		case "quit":
			fmt.Println("Thanks for banking with us")
			return
		default:
			fmt.Println("Invalid Input: Type help")
		}

		for {
			fmt.Println("Type ok to continue or quit to exit")
			_, err = fmt.Scan(&cont)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			cont = strings.Trim(cont, "'\"")
			if cont == "ok" {
				break
			}
			if cont == "quit" {
				fmt.Println("Thanks for banking with us")
				return
			}
			fmt.Println("please type 'ok' or 'quit'")
		}
	}
}
