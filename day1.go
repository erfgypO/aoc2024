package aoc24

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input/day1.txt
var d1input string

func parseDay1Input() ([]int, []int, error) {
	var left []int
	var right []int

	for _, line := range strings.Split(d1input, "\n") {
		if line == "" {
			continue
		}

		splits := strings.Split(line, "   ")

		i1, err := strconv.Atoi(splits[0])
		if err != nil {
			return left, right, err
		}

		i2, err := strconv.Atoi(splits[1])
		if err != nil {
			return left, right, err
		}

		left = append(left, i1)
		right = append(right, i2)
	}

	return left, right, nil
}

func (t *Aoc24) D1P1() {
	left, right, err := parseDay1Input()
	if err != nil {
		log.Fatal(err)
	}
	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0

	for i := range left {
		leftLowest := left[i]
		rightLowest := right[i]

		distance := leftLowest - rightLowest
		if distance < 0 {
			distance = distance * -1
		}

		totalDistance += distance
	}

	log.Println(fmt.Sprintf("Total distance: %d", totalDistance))
}

func (t *Aoc24) D1P2() {
	left, right, err := parseDay1Input()
	if err != nil {
		log.Fatal(err)
	}

	similarity := 0

	for _, leftValue := range left {
		localSimilarity := 0
		for _, rightValue := range right {
			if leftValue == rightValue {
				localSimilarity++
			}
		}

		similarity += leftValue * localSimilarity
	}

	log.Println(fmt.Sprintf("Similarity: %d", similarity))
}
