package main

type routeCategoryTraf struct {
	AbbreviatedNumberTrafficCases            string `json:"abbreviatedNumberTrafficCases"`
	TCDCategoryNightForIncomingExternalLines string `json:"tcdCategoryNightForIncomingExternalLines"`
	TCDCategoryDayForIncomingExternalLines   string `json:"tcdCategoryDayForIncomingExternalLines"`
	TrafficConnectionClass                   string `json:"trafficConnectionClass"`
}

func createNewRouteCategoryTraf(data []rune) routeCategoryTraf {
	return routeCategoryTraf{
		AbbreviatedNumberTrafficCases:            string(data[0:2]),
		TCDCategoryNightForIncomingExternalLines: string(data[2:4]),
		TCDCategoryDayForIncomingExternalLines:   string(data[4:6]),
		TrafficConnectionClass:                   string(data[6:8]),
	}
}
