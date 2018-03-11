package main

import (
	"io/ioutil"
	"log"
)

type routeCategoryData struct {
	Route        string            `json:"route"`
	Cust         string            `json:"cust"`
	Selection    routeCategorySel  `json:"selection"`
	Service      routeCategoryServ `json:"service"`
	Trm          string            `json:"trm"`
	Nodg         string            `json:"nodg"`
	Dist         string            `json:"dist"`
	Disl         string            `json:"disl"`
	Traffic      routeCategoryTraf `json:"traffic"`
	serviceRaw   string
	trafficRaw   string
	sigRaw       string
	bcapRaw      string
	selectionRaw string
}

func (routeCategory *routeCategoryData) ReadFromFile() []routeCategoryData {
	data, err := ioutil.ReadFile("pbx_data/ROCAP")
	if err != nil {
		log.Fatal(err)
	}

	return routeCategory.parse(data)
}

func (routeCategory *routeCategoryData) parse(data []byte) []routeCategoryData {
	var result []routeCategoryData

	values := getColumnValues(string(data), 2)
	for _, v := range values {
		routeCat := routeCategoryData{
			Route:        string(v[0]),
			Cust:         string(v[1]),
			selectionRaw: string(v[2]),
			Trm:          string(v[3]),
			serviceRaw:   string(v[4]),
			Nodg:         string(v[5]),
			Dist:         string(v[6]),
			Disl:         string(v[7]),
			trafficRaw:   string(v[8]),
			sigRaw:       string(v[9]),
			bcapRaw:      string(v[10]),
		}

		routeCat.Selection = createNewRouteCategorySel([]rune(routeCat.selectionRaw))
		routeCat.Service = createNewRouteCategoryServ([]rune(routeCat.serviceRaw))
		routeCat.Traffic = createNewRouteCategoryTraf([]rune(routeCat.trafficRaw))

		result = append(result, routeCat)
	}

	return result
}
