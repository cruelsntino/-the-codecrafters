// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Wealth Ibe]
// Squad:  [The Structs]

package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"
)

func upper(s string) string {
	if len(s) == 0 {
		fmt.Println("Empty string")
	}
	return strings.ToUpper(s + " ")
}

func lower(s string) string {
	if len(s) == 0 {
		fmt.Println("Empty string")
	}
	return strings.ToLower(s)
}

func cap(s string) string {
	if len(s) == 0 {
		fmt.Println("No input")
	}
	txt := strings.Fields(s)
	for i, v := range txt {
		txt[i] = strings.ToUpper(string(v[0])) + strings.ToLower(v[1:])
	}
	return strings.Join(txt, " ")
}
func title(s string) string {
	if len(s) == 0 {
		fmt.Println("No input")
	}
	txt := strings.Fields(s)
	small := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "but": true, "or": true,
		"for": true, "nor": true,
		"on": true, "at": true, "to": true,
		"by": true, "in": true, "of": true,
		"up": true, "as": true, "is": true, "it": true,
	}
	for i, char := range txt {
		line := strings.ToLower(char)
		if i == 0 {
			txt[i] = strings.ToUpper(string(line[0])) + line[1:]
			continue
		}
		if small[line] {
			txt[i] = line
		} else {
			txt[i] = strings.ToUpper(string(line[0])) + line[1:]
		}
	}
	return strings.Join(txt, " ")
}

func reverse(s string) string {
	words := strings.Fields(s)
	for i, char := range words {
		runes := []rune(char)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")

}

func snake(s string) string {
	return strings.ToLower(strings.Join(strings.Fields(s), "_"))
}

func help() {
	fmt.Println("SUPPORTED COMMANDS:")
	fmt.Printf("\nupper <text> -> Convert every letter to UPPERCASE\n")
	fmt.Printf("\nlower <text> -> Convert every letter to lowercase\n")
	fmt.Printf("\ncap <text> -> Capitalise the first letter of every word\n")
	fmt.Printf("\ntitle <text> -> Title Case — like cap, but small connector words stay lowercase unless they are the first word\n")
	fmt.Printf("\nsnake <text> -> Convert to snake_case. All lowercase, spaces replaced with _.\n")
	fmt.Printf("\nreverse <text> -> Reverse each word individually. Word order stays the same. Spaces between words are preserved\n")
	fmt.Printf("\ntype 'exit' to quit\n\n")
}

func main() {
	fmt.Println("SENTINEL STRING TRANSFORMER - ONLINE")
	fmt.Println("Enter 'help' for guidelines or 'exit' to quit")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println("Error: no input")
			continue
		}
		if strings.ToLower(input) == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}
		if input == "help" {
			help()
			continue
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])
		parts[0] = strings.ToLower(command)
		if parts[0] != "help" && parts[0] != "upper" && parts[0] != "lower" && parts[0] != "cap" && parts[0] != "title" && parts[0] != "snake" && parts[0] != "reverse" && parts[0] != "exit" {
			fmt.Printf("Unknown Command: '%s'\n", command)
			fmt.Printf("Valid commands:\n upper\n lower\n cap\n title\n snake\n reverse\n exit\n")
			continue
		}
		if len(parts) == 1 {
			fmt.Printf("No text provided. Usage: %s <text>\n", command)
			continue
		}
		text := strings.Join(parts[1:], " ")
		switch command {
		case "upper":
			fmt.Println(upper(text))
		case "lower":
			fmt.Println(lower(text))
		case "cap":
			fmt.Println(cap(text))
		case "title":
			fmt.Println(title(text))
		case "snake":
			// result := strings.Fields(text)
			fmt.Println(snake(text))
		case "reverse":
			fmt.Println(reverse(text))
		default:
			fmt.Println("Invalid command")
		}
	}
}
