package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// asd := NumberConversionEntry{}
	// lol := asd.ReadFromFile()

	// brap, err := json.MarshalIndent(lol, "", "    ")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(brap))

	asd := AcdGroupCallData{}

	brap, err := json.MarshalIndent(asd.ReadFromFile(), "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(brap))
}
