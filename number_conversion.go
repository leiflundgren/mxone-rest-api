package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// NumberConversionEntry is used for printing data from the number conversion and bearer Capability and High-Level Compatibility substitution tables.
type NumberConversionEntry struct {
	conversionType    int
	entry             string
	numberType        int
	pre               string
	route             int
	targetDestination string
}

// NumberConversionEntries is an array of Number Conversion entries
type NumberConversionEntries []NumberConversionEntry

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

func (this *NumberConversionEntries) Get() {

}

func (this *NumberConversionEntries) execAndParse() {
	cmd := exec.Command("number_conversion_print")

	var bytesResult bytes.Buffer
	cmd.Stdout = &bytesResult

	if cmd.Run() != nil {
		log.Fatal("Nu gick något åt helvete")
	}

	fmt.Println(bytesResult.String())
}
