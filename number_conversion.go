package main

import (
	"io/ioutil"
	"log"
)

type numberConversionEntry struct {
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

func (numberConversion numberConversionEntry) ReadFromFile() []numberConversionEntry {
	byteArray, err := ioutil.ReadFile("pbx_data/number_conversion_print")
	if err != nil {
		log.Println(err)
	}

	return numberConversion.parse(byteArray)
}

func (numberConversion numberConversionEntry) parse(data []byte) []numberConversionEntry {
	var result []numberConversionEntry

	values := getColumnValues(string(data), 2)
	for _, v := range values {
		numberConv := numberConversionEntry{
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
