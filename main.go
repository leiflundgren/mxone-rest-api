package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	asd := RouteCategoryData{}
	brap, _ := json.MarshalIndent(asd.ReadFromFile(), "", "    ")
	fmt.Println(string(brap))

	qwe := NumberConversionEntry{}
	brap, _ = json.MarshalIndent(qwe.ReadFromFile(), "", "    ")

	fmt.Println(string(brap))
}
