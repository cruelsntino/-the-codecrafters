package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	green := "\033[1;92m"
	const (
		red = "\033[91m"

		blue  = "\033[34m"
		reset = "\033[0m"
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(green + "==================BASE CONVERTER==================" + reset)
	fmt.Printf("")

	for {
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)

		if input == "quit" {
			fmt.Println(green + "Thanks for Banking with us!" + reset)
			return
		}
		if input == "help" {
			fmt.Println(blue + "Please follow this guidelines:" + reset)
			fmt.Println(blue + "convert 1E hex" + reset)
			fmt.Println(blue + "convert 10101 bin" + reset)
			fmt.Println(blue + "convert 1285 dec" + reset)

			continue
		}
		if len(input) == 0 {
			fmt.Println(red + "ERROR: Cannot be empty" + reset)
			continue
		}
		split := strings.Fields(input)
		if len(split) < 3 {
			fmt.Println("Not enough arguments.")
		}
		if len(split) > 3 {
			fmt.Println("Too many Arguments")
		}
		command := "convert"
		if command != split[0] {
			fmt.Println("Invalid Syntax, enter help for guidelines:")
			continue
		}
		value := split[1]
		base := split[2]

		switch base {
		case "hex":
			dec, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Printf("ERROR: '%s' is an invalid hex\n", value)
				continue
			}
			fmt.Printf("Decimal: %d\n", dec)
		case "dec":
			dec, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Printf("ERROR: '%s' is an invalid dec\n", value)
				continue
			}
			fmt.Printf("Binary: %b\n", dec)
			fmt.Printf("Hexadecimal: %X\n", dec)
		case "bin":
			dec, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Printf("ERROR: '%s' is an invalid bin\n", value)
				continue
			}
			fmt.Printf("Decimal: %d", dec)
		default:
			fmt.Println("Inavlid Base: try again")
		}

	}
}
