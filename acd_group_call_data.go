package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
)

// AcdGroupCallData - Automatic Call Distribution Group Category
type AcdGroupCallData struct {
	Group   string `json:"group"`
	Lim     string `json:"lim"`
	servRaw string
	Serv    AcdGroupServ `json:"serv"`
	Traf    string       `json:"traf"`
	selRaw  string
	Sel     AcdGroupSel `json:"sel"`
	Que     AcdGroupQue `json:"que"`
	queRaw  string
	Cust    string `json:"cust"`
	Sat     string `json:"sat"`
	Tpcs    string `json:"tpcs"`
	Actc    string `json:"actc"`
}

// ReadFromFile Allows reading from a PC-Regen file or other data dump
func (acdGroup *AcdGroupCallData) ReadFromFile() []AcdGroupCallData {
	byteArray, err := ioutil.ReadFile("pbx_data/ACGCP")
	if err != nil {
		log.Fatal(err)
	}

	return acdGroup.parse(byteArray)
}

func (acdGroup *AcdGroupCallData) parse(byteArray []byte) []AcdGroupCallData {
	var linesToSkip = 1
	var skippedLines = 0
	var result []AcdGroupCallData

	scanner := bufio.NewScanner(strings.NewReader(string(byteArray)))
	for scanner.Scan() {
		// Skipping the table header
		if skippedLines <= linesToSkip {
			skippedLines++
			continue
		}

		if len(scanner.Text()) == 0 || len(scanner.Text()) <= 3 {
			continue
		}

		runes := []rune(scanner.Text())
		runesLength := len(runes)

		var grp string
		if runesLength >= 13 {
			grp = string(runes[0:13])
			grp = strings.TrimSpace(grp)
		}

		var lim string
		if runesLength >= 20 {
			lim = string(runes[13:20])
			lim = strings.TrimSpace(lim)
		}

		var serv string
		if runesLength >= 28 {
			serv = string(runes[20:28])
			serv = strings.TrimSpace(serv)
		}

		var traf string
		if runesLength >= 35 {
			traf = string(runes[28:35])
			traf = strings.TrimSpace(traf)
		}

		var sel string
		if runesLength >= 41 {
			sel = string(runes[35:41])
			sel = strings.TrimSpace(sel)
		}

		var que string
		if runesLength >= 55 {
			que = string(runes[42:55])
			que = strings.TrimSpace(que)
		}

		var cust string
		if runesLength >= 61 {
			cust = string(runes[55:61])
			cust = strings.TrimSpace(cust)
		}

		var sat string
		if runesLength >= 66 {
			sat = string(runes[61:66])
			sat = strings.TrimSpace(sat)
		}

		var tpcs string
		if runesLength >= 72 {
			tpcs = string(runes[66:72])
			tpcs = strings.TrimSpace(tpcs)
		}

		var actc string
		if runesLength >= 77 {
			actc = string(runes[72:77])
			actc = strings.TrimSpace(actc)
		}

		queStruct := CreateNewAcdGroupQue(que)
		servStruct := CreateNewAcdGroupServ(serv)

		acdGroupEntry := AcdGroupCallData{
			Group:   grp,
			Lim:     lim,
			servRaw: serv,
			Serv:    servStruct,
			Traf:    traf,
			selRaw:  sel,
			queRaw:  que,
			Que:     queStruct,
			Cust:    cust,
			Sat:     sat,
			Tpcs:    tpcs,
			Actc:    actc,
		}

		result = append(result, acdGroupEntry)
	}

	return result
}
