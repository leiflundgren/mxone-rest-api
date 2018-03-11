package main

import (
	"bufio"
	"regexp"
	"strings"
)

func getColumnsWidth(data string, tableHeaderLine int) []int {
	var result []int
	var readLines int

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		readLines++

		if readLines == tableHeaderLine {
			reg := regexp.MustCompile(`(\w+\s+)`)

			for _, match := range reg.FindAllString(scanner.Text(), -1) {
				result = append(result, len(match))
			}
		}
	}

	return result
}
