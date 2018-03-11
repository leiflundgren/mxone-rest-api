package main

import (
	"io/ioutil"
	"log"
)

type RouteCategoryData struct {
	Rou       string `json:"route"`
	Cust      string `json:"cust"`
	selRaw    string
	Selection routeCategoryDataSel `json:"selection"`
	Trm       string               `json:"trm"`
	servRaw   string
	Nodg      string `json:"nodg"`
	Dist      string `json:"dist"`
	Disl      string `json:"disl"`
	trafRaw   string
	sigRaw    string
	bcapRaw   string
}

func (routeCategory *RouteCategoryData) ReadFromFile() []RouteCategoryData {
	data, err := ioutil.ReadFile("pbx_data/ROCAP")
	if err != nil {
		log.Println(err)
		return nil
	}

	return routeCategory.parse(data)
}

func (routeCategory *RouteCategoryData) parse(data []byte) []RouteCategoryData {
	var result []RouteCategoryData

	values := getColumnValues(string(data), 2)
	for _, v := range values {
		routeCat := RouteCategoryData{
			Rou:     string(v[0]),
			Cust:    string(v[1]),
			selRaw:  string(v[2]),
			Trm:     string(v[3]),
			servRaw: string(v[4]),
			Nodg:    string(v[5]),
			Dist:    string(v[6]),
			Disl:    string(v[7]),
			trafRaw: string(v[8]),
			sigRaw:  string(v[9]),
			bcapRaw: string(v[10]),
		}

		routeCat.Selection = createNewRouteCategorySel(routeCat.selRaw)

		result = append(result, routeCat)
	}

	return result
}
