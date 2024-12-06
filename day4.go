package aoc24

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input/day4.txt
var d4Input string

type D4Data struct {
	matrix [][]string
	width  int
	height int
}

func parseInput() *D4Data {
	lines := strings.Split(d4Input, "\n")

	width := len(lines[0])
	height := len(lines)

	data := D4Data{
		matrix: make([][]string, height),
		width:  width,
		height: height,
	}

	for y := range data.matrix {
		data.matrix[y] = make([]string, width)
	}

	for y, line := range lines {
		for x, char := range line {
			data.matrix[x][y] = string(char)
		}
	}

	return &data
}

func findWord(matrix [][]string, x, y, dx, dy int, word string) bool {
	for i := 0; i < len(word); i++ {
		nx, ny := x+i*dx, y+i*dy
		if nx < 0 || ny < 0 || ny >= len(matrix) || nx >= len(matrix[0]) {
			return false
		}
		if matrix[ny][nx] != string(word[i]) {
			return false
		}
	}
	return true
}

func countWords(matrix [][]string, word string) int {
	directions := [][2]int{{1, 0}, {0, 1}, {1, 1}, {-1, 1}} // Right, Down, Diagonal down-right, Diagonal up-right
	count := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if findWord(matrix, x, y, dx, dy, word) {
					count++
				}
			}
		}
	}

	return count
}

func (t *Aoc24) D4P1() {
	wordCount := 0
	input := parseInput()

	for y := range input.height {
		for x := range input.width {
			if x+3 < input.width {
				word := fmt.Sprintf("%s%s%s%s", input.matrix[x][y], input.matrix[x+1][y], input.matrix[x+2][y], input.matrix[x+3][y])
				if word == "XMAS" || word == "SAMX" {
					wordCount++
				}
			}

			if y+3 < input.height {
				word := fmt.Sprintf("%s%s%s%s", input.matrix[x][y], input.matrix[x][y+1], input.matrix[x][y+2], input.matrix[x][y+3])
				if word == "XMAS" || word == "SAMX" {
					wordCount++
				}
			}

			if y+3 < input.height && x+3 < input.width {
				word := fmt.Sprintf("%s%s%s%s", input.matrix[x][y], input.matrix[x+1][y+1], input.matrix[x+2][y+2], input.matrix[x+3][y+3])
				if word == "XMAS" || word == "SAMX" {
					wordCount++
				}
			}

			if y-3 >= 0 && x-3 >= 0 {
				word := fmt.Sprintf("%s%s%s%s", input.matrix[x][y], input.matrix[x-1][y-1], input.matrix[x-2][y-2], input.matrix[x-3][y-3])
				if word == "XMAS" || word == "SAMX" {
					wordCount++
				}
			}
		}
	}

	log.Printf("Count: %d", wordCount)

	gptResult := countWords(input.matrix, "XMAS") + countWords(input.matrix, "SAMX")
	log.Printf("GptResult: %d", gptResult)
}

func isXMASPattern(matrix [][]string, x, y int) bool {
	// Check if (x, y) can form an "X-MAS" pattern
	if y-1 < 0 || y+1 >= len(matrix) || x-1 < 0 || x+1 >= len(matrix[0]) {
		return false // Out of bounds
	}

	// Extract the two diagonals
	topLeft := fmt.Sprintf("%s%s%s", matrix[y-1][x-1], matrix[y][x], matrix[y+1][x+1])  // `/`
	topRight := fmt.Sprintf("%s%s%s", matrix[y-1][x+1], matrix[y][x], matrix[y+1][x-1]) // `\`

	// Check if both diagonals form "MAS" or "SAM"
	isValid := (topLeft == "MAS" || topLeft == "SAM") && (topRight == "MAS" || topRight == "SAM")
	return isValid
}

func countXMASPatterns(matrix [][]string) int {
	count := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isXMASPattern(matrix, x, y) {
				count++
			}
		}
	}

	return count
}

func (t *Aoc24) D4P2() {
	input := parseInput()
	matrix := input.matrix

	count := countXMASPatterns(matrix)
	log.Printf("Count of X-MAS patterns: %d", count)
}
