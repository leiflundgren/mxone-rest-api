package main

import (
	"io/ioutil"
	"log"
)

// NumberConversionEntry is used for printing data from the number conversion and bearer Capability and High-Level Compatibility substitution tables.
type NumberConversionEntry struct {
	Entry             string `json:"entry"`
	ConversionType    string `json:"conversionType"`
	NumberType        string `json:"numberType"`
	TargetDestination string `json:"targetDestination"`
	Route             string `json:"route"`
	Pre               string `json:"pre"`
	Trc               string `json:"trc"`
	NewTyp            string `json:"newTyp"`
	Cont              string `json:"const"`
	Bcap              string `json:"bcap"`
	Hlc               string `json:"hlc"`
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
		log.Println(err)
	}

	return numberConversion.parse(byteArray)
}

func (numberConversion *NumberConversionEntry) parse(data []byte) []NumberConversionEntry {
	var result []NumberConversionEntry

	values := getColumnValues(string(data), 2)
	for _, v := range values {
		numberConv := NumberConversionEntry{
			Entry:             string(v[0]),
			ConversionType:    string(v[1]),
			NumberType:        string(v[2]),
			TargetDestination: string(v[3]),
			Route:             string(v[4]),
			Pre:               string(v[5]),
			Trc:               string(v[6]),
			NewTyp:            string(v[7]),
			Cont:              string(v[8]),
			Bcap:              string(v[9]),
			Hlc:               string(v[10]),
		}

		result = append(result, numberConv)
	}

	return result
}
