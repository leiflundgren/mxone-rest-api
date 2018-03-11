package main

import (
	"io/ioutil"
	"log"
)

type acdGroupCallData struct {
	Group        string       `json:"group"`
	Lim          string       `json:"lim"`
	Service      acdGroupServ `json:"service"`
	Traf         string       `json:"traf"`
	Selection    acdGroupSel  `json:"sel"`
	Queue        acdGroupQue  `json:"que"`
	Cust         string       `json:"cust"`
	Sat          string       `json:"sat"`
	Tpcs         string       `json:"tpcs"`
	Actc         string       `json:"actc"`
	queueRaw     string
	serviceRaw   string
	selectionRaw string
	trafficRaw   string
}

func (acdGroup acdGroupCallData) ReadFromFile() []acdGroupCallData {
	byteArray, err := ioutil.ReadFile("pbx_data/ACGCP")
	if err != nil {
		log.Fatal(err)
	}

	return acdGroup.parse(byteArray)
}

func (acdGroup acdGroupCallData) parse(data []byte) []acdGroupCallData {
	var result []acdGroupCallData

	values := getColumnValues(string(data), 2)
	for _, v := range values {
		acdGroupEntry := acdGroupCallData{
			Group:        v[0],
			Lim:          v[1],
			serviceRaw:   v[2],
			trafficRaw:   v[3],
			selectionRaw: v[4],
			queueRaw:     v[5],
			Cust:         v[6],
			Sat:          v[7],
			Tpcs:         v[8],
			Actc:         v[9],
		}

		acdGroupEntry.Queue = createNewAcdGroupQue([]rune(acdGroupEntry.queueRaw))
		acdGroupEntry.Service = createNewAcdGroupServ([]rune(acdGroupEntry.serviceRaw))
		acdGroupEntry.Selection = createNewAcdGroupSel([]rune(acdGroupEntry.selectionRaw))

		result = append(result, acdGroupEntry)
	}

	return result
}
