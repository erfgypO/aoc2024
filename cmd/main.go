package main

import (
	"bufio"
	"fmt"
	"github.com/erfgypO/aoc24"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent of Code 2024")
	fmt.Print("Select your day and part (eg. D1P1): ")
	input, _ := reader.ReadString('\n')

	input = strings.Trim(strings.Replace(input, "\n", "", -1), " ")
	aoc24.Run(input)
}
