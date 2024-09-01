package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	af "asciiart/functions"
)

func main() {
	input := os.Args[1]
	if af.ContainsOnly(input) {
		for i := 0; i < len(input)/2; i++ {
			fmt.Println()
		}
		return

	}

	inputsplit := strings.Split(input, "\\n")
	for _, line := range inputsplit {
		for _, c := range line {
			if c < 32 || c > 126 {
				log.Fatal("Error : input should only contain printable ascii characters")
			}
		}
	}
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal("Error : couldn't read file")
	}
	Lines := strings.Split(string(content), "\n")

	for _, line := range inputsplit {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				fmt.Print(Lines[(int(char)-31)*9-(8-i)])
			}
			fmt.Println()
		}
	}
}
