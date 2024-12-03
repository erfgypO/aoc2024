package aoc24

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input/day3.txt
var d3input string

func (t *Aoc24) D3P1() {
	var re = regexp.MustCompile(`(?m)mul\(\d+,\d+\)`)

	total := 0

	for _, match := range re.FindAllString(d3input, -1) {
		total += evaluateMulFunction(match)
	}

	log.Printf("Total values: %d", total)
}

func (t *Aoc24) D3P2() {
	total := 0
	enabled := true

	data := d3input
	i := 0
	for i < len(data) {
		if i+4 <= len(data) && data[i:i+4] == "do()" {
			enabled = true
			i += 4
			continue
		}

		if i+7 <= len(data) && data[i:i+7] == "don't()" {
			enabled = false
			i += 7
			continue
		}

		if i+4 <= len(data) && data[i:i+4] == "mul(" && enabled {
			i += 4
			num1 := ""

			for i < len(data) && unicode.IsDigit(rune(data[i])) {
				num1 += string(data[i])
				i++
			}

			if i < len(data) && data[i] == ',' {
				i++
				num2 := ""

				for i < len(data) && unicode.IsDigit(rune(data[i])) {
					num2 += string(data[i])
					i++
				}

				if i < len(data) && data[i] == ')' && num1 != "" && num2 != "" {
					n1, _ := strconv.Atoi(num1)
					n2, _ := strconv.Atoi(num2)
					total += n1 * n2
				}
			}
		}

		i++
	}

	log.Printf("Total values: %d", total)
}

func evaluateMulFunction(mulFunctionString string) int {
	var values []int

	for _, split := range strings.Split(mulFunctionString, ",") {
		sanitizedValue := split

		if strings.HasPrefix(sanitizedValue, "mul(") {
			sanitizedValue = sanitizedValue[4:]
		} else {
			sanitizedValue = strings.TrimRight(sanitizedValue, ")")
		}

		number, err := strconv.Atoi(sanitizedValue)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, number)
	}

	return values[0] * values[1]
}
