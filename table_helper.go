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

			// HACK, appending whitespaces due to poor regex
			tableHeader := scanner.Text() + "              "

			for _, match := range reg.FindAllString(tableHeader, -1) {
				result = append(result, len(match))
			}
		}
	}

	return result
}

func getColumnValues(data string, tableHeaderLine int) map[int][]string {
	var readLines int
	var lineNumber = -1
	var columns = getColumnsWidth(string(data), 2)
	result := make(map[int][]string)

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		lineNumber++

		if readLines < 2 || len(scanner.Text()) == 0 || scanner.Text() == "" || len(scanner.Text()) < 4 {
			readLines++
			continue
		}

		runes := []rune(scanner.Text())
		runesLength := len(runes)
		var readBytes int

		for i, v := range columns {
			if readBytes == 0 {
				// First column
				result[lineNumber] = append(result[lineNumber], strings.TrimSpace(string(runes[0:v])))
			} else if i == len(columns)-1 {
				// Last column
				result[lineNumber] = append(result[lineNumber], strings.TrimSpace(string(runes[readBytes:runesLength])))
			} else {
				// Everything in between
				result[lineNumber] = append(result[lineNumber], strings.TrimSpace(string(runes[readBytes:readBytes+v])))
			}

			readBytes += v
		}
	}

	return result
}
