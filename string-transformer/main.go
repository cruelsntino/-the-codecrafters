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
	return strings.Join(txt, "")
}
func title(s string) string {
	if len(s) == 0 {
		fmt.Println("No input")
	}
	return strings.Title(s)

}

func reverse(s string) string {
	txt := ""
	for i := len(s) - 1; i >= 0; i-- {
		txt += string(s[i])
	}
	return txt
}

func snake(s []string) string {
	txt := ""
	txt1 := ""
	for i := 0; i < len(s); i++ {
		txt += s[i] + ""
	}
	for j := 0; j < len(txt)-1; j++ {
		txt1 += string(txt[j])
	}
	txt1 = strings.ToLower(txt1)
	txt1 = strings.ReplaceAll(txt1, "", "_")
	return txt1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println("Error: no input")
			continue
		}
		if strings.ToLower(input) == "exit" {
			fmt.Println("Thanks for banking with us.")
			return
		}
		split := strings.Fields(input)
		command := strings.ToLower(split[0])
		if len(split) == 1 {
			fmt.Printf("usage: %s <text>\n", command)
			continue
		}
		text := strings.Join(split[1:], "")
		switch command {
		case "1":
			fmt.Println(upper(text))
		case "2":
			fmt.Println(lower(text))
		case "3":
			fmt.Println(cap(text))
		case "4":
			fmt.Println(title(text))
			// case "5":
			// 	result := strings.ToLower(text)
			// 	fmt.Println(snake(result))
		}
	}
}
