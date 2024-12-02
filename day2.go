package aoc24

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input/day2.txt
var d2input string

func parseDay2Input() ([][]int, error) {
	var lines [][]int

	for _, line := range strings.Split(d2input, "\n") {
		if line == "" {
			continue
		}

		var lineData []int
		splits := strings.Split(line, " ")
		for _, split := range splits {
			number, err := strconv.Atoi(split)
			if err != nil {
				return nil, err
			}
			lineData = append(lineData, number)
		}
		lines = append(lines, lineData)
	}
	return lines, nil
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func isSafeReport(report []int) bool {

	reversedReport := slices.Clone(report)
	slices.Reverse(reversedReport)

	sorted := slices.IsSorted(report) || slices.IsSorted(reversedReport)

	var changes []int

	for index, value := range report {
		if index+1 == len(report) {
			continue
		}
		change := value - report[index+1]
		changes = append(changes, change)
	}

	invalidChangeStep := false

	for _, change := range changes {
		absoluteChange := abs(change)

		if absoluteChange == 0 || absoluteChange > 3 {
			invalidChangeStep = true
		}
	}

	safeReport := sorted && !invalidChangeStep
	return safeReport
}

func isToleratedReport(report []int) bool {
	reversedReport := slices.Clone(report)
	slices.Reverse(reversedReport)

	sorted := slices.IsSorted(report) || slices.IsSorted(reversedReport)

	var changes []int

	for index, value := range report {
		if index+1 == len(report) {
			continue
		}
		change := value - report[index+1]
		changes = append(changes, change)
	}

	unsafeChanges := 0

	for _, change := range changes {
		absoluteChange := abs(change)
		if absoluteChange == 0 || absoluteChange > 3 {
			unsafeChanges++
		}
	}

	dampedReport := false

	if !sorted || unsafeChanges == 1 || unsafeChanges == 2 {
		for key := range report {
			modifiedReport := slices.Clone(report)
			modifiedReport = append(modifiedReport[:key], modifiedReport[key+1:]...)
			dampedReport = isSafeReport(modifiedReport)
			if dampedReport {
				break
			}
		}
	}

	toleratedReport := (sorted && unsafeChanges == 0) || dampedReport
	return toleratedReport
}

func (t *Aoc24) D2P1() {
	lines, err := parseDay2Input()
	if err != nil {
		log.Panicf("Error parsing day2 input: %v", err)
	}

	safeReports := 0

	for _, l := range lines {

		if isSafeReport(l) {
			safeReports++
		}
	}

	log.Printf("SafeReports: %d", safeReports)
}

func (t *Aoc24) D2P2() {
	lines, err := parseDay2Input()
	if err != nil {
		log.Panicf("Error parsing day2 input: %v", err)
	}

	toleratedReports := 0

	for _, l := range lines {

		if isToleratedReport(l) {
			toleratedReports++
		}
	}

	log.Printf("ToleratedReports: %d", toleratedReports)
}
