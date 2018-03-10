package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// NumberConversionEntry is used for printing data from the number conversion and bearer Capability and High-Level Compatibility substitution tables.
type NumberConversionEntry struct {
	ConversionType    int    `json:"conversionType"`
	Entry             string `json:"entry"`
	NumberType        int    `json:"numberType"`
	Pre               string `json:"pre"`
	Route             int    `json:"route"`
	TargetDestination string `json:"targetDestination"`
}

// Conversion Type
const (
	receivedBNumber                              = 0
	sentANumberAndConnectedNumber                = 1
	receivedANumberAndReceivedConnectedNumber    = 2
	bearerCapability                             = 3
	inboundConversion                            = 4
	sentANumberAndSentConnectedNumberToExtension = 5
	receivedANumberForCallsFromMobile            = 6
	receivedPublicANumber                        = 7
)

// Number Type
const (
	unknownPublicNumber                   = 0
	internationalNumber                   = 1
	nationalNumber                        = 2
	networkSpecificNumber                 = 3
	localPublicNumber                     = 4
	unknownPrivateNumber                  = 5
	localPrivateNumber                    = 6
	levelOneRegionalNumber                = 7
	internalDirectoryNumberPublicNetwork  = 10
	internalDirectoryNumberPrivateNetowkr = 11
	internalDirectoryContinue             = 12
)

// ReadFromFile Allows reading from a PC-Regen file or other data dump
func (numberConversion *NumberConversionEntry) ReadFromFile() []NumberConversionEntry {
	byteArray, err := ioutil.ReadFile("number_conversion_print")
	if err != nil {
		log.Fatal(err)
	}

	return numberConversion.parse(byteArray)
}

func (numberConversion *NumberConversionEntry) parse(byteArray []byte) []NumberConversionEntry {
	var tempVar string
	var linesToSkip = 2
	var skippedLines = 0
	var result []NumberConversionEntry

	scanner := bufio.NewScanner(strings.NewReader(string(byteArray)))
	for scanner.Scan() {
		// Skipping the table header
		if skippedLines <= linesToSkip {
			skippedLines++
			continue
		}

		fmt.Println(scanner.Text())
		runes := []rune(scanner.Text())
		runesLength := len(runes)

		var entry string
		if runesLength >= 20 {
			entry = string(runes[0:20])
			entry = strings.TrimSpace(entry)
		}

		var conversionType int
		if runesLength >= 22 {
			tempVar = string(runes[21:22])
			tempVar = strings.TrimSpace(tempVar)

			conversionType, _ = strconv.Atoi(tempVar)
		}

		var numberType int
		if runesLength >= 29 {
			tempVar = string(runes[28:29])
			tempVar = strings.TrimSpace(tempVar)

			numberType, _ = strconv.Atoi(tempVar)
		}

		var route int
		if runesLength >= 36 {
			tempVar = string(runes[35:39])
			tempVar = strings.TrimSpace(tempVar)

			route, _ = strconv.Atoi(tempVar)
		}

		var targetDestination string
		if runesLength >= 40 {
			targetDestination = string(runes[39:47])
			targetDestination = strings.TrimSpace(targetDestination)
		}

		var pre string
		if runesLength >= 48 {
			pre = string(runes[47:68])
			pre = strings.TrimSpace(pre)
		}

		numberConversionEntry := NumberConversionEntry{}
		numberConversionEntry.ConversionType = conversionType
		numberConversionEntry.Entry = entry
		numberConversionEntry.Pre = pre
		numberConversionEntry.TargetDestination = targetDestination
		numberConversionEntry.Route = route
		numberConversionEntry.NumberType = numberType

		result = append(result, numberConversionEntry)
	}

	return result
}

func (numberConversion *NumberConversionEntry) execAndParse() {
	cmd := exec.Command("number_conversion_print")

	var bytesResult bytes.Buffer
	cmd.Stdout = &bytesResult

	if cmd.Run() != nil {
		log.Fatal("Nu gick något åt helvete")
	}

	fmt.Println(bytesResult.String())
	scanner := bufio.NewScanner(strings.NewReader(bytesResult.String()))
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		fmt.Println(string(runes[0:21]))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Fan också")
	}
}
