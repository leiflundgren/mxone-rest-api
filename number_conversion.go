package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// NumberConversionEntry is used for printing data from the number conversion and bearer Capability and High-Level Compatibility substitution tables.
type NumberConversionEntry struct {
	ConversionType    string `json:"conversionType"`
	Entry             string `json:"entry"`
	NumberType        string `json:"numberType"`
	Pre               string `json:"pre"`
	Route             string `json:"route"`
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
	byteArray, err := ioutil.ReadFile("pbx_data/number_conversion_print")
	if err != nil {
		log.Fatal(err)
	}

	return numberConversion.parse(byteArray)
}

func (numberConversion *NumberConversionEntry) parse(data []byte) []NumberConversionEntry {
	var result []NumberConversionEntry

	values := getColumnValues(string(data), 2)
	for _, v := range values {
		numberConv := NumberConversionEntry{
			Entry: string(v[0]),
		}

		result = append(result, numberConv)
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
