package main

type routeCategoryDataSel struct {
	DirectIndialingCharacteristics             string `json:"direct-indialing-characteristics"`
	IncomingTraffic                            string `json:"incoming-traffic"`
	LineSelection                              string `json:"line-selection"`
	RouteCharacteristicsOutgoingTraffic        string `json:"route-characteristics-outgoing-traffic"`
	AlternativeRouteSelectionOnIncomingTraffic string `json:"alternative-route-selection-on-incoming-traffic"`
	CustomerAffiliation                        string `json:"customer-affiliation"`
	VirtualCalls                               string `json:"virtual-calls"`
	MaliciousCallTracing                       string `json:"malicious-call-tracing"`
	FRLCategoryForIncoming                     string `json:"frl-category-for-incoming"`
	CallServiceInformation                     string `json:"call-service-information"`
	ReceivedTCM                                string `json:"received-tcm"`
	TollExchangeCategory                       string `json:"toll-exchange-category"`
	RouteToTelidentMachine                     string `json:"route-to-telident-machine"`
}

func createNewRouteCategorySel(data string) routeCategoryDataSel {
	runes := []rune(data)

	return routeCategoryDataSel{
		DirectIndialingCharacteristics:             string(runes[0:1]),
		IncomingTraffic:                            string(runes[1:2]),
		LineSelection:                              string(runes[2:3]),
		RouteCharacteristicsOutgoingTraffic:        string(runes[3:4]),
		AlternativeRouteSelectionOnIncomingTraffic: string(runes[4:5]),
		CustomerAffiliation:                        string(runes[5:8]),
		VirtualCalls:                               string(runes[8:9]),
		MaliciousCallTracing:                       string(runes[9:10]),
		FRLCategoryForIncoming:                     string(runes[10:11]),
		CallServiceInformation:                     string(runes[11:12]),
		ReceivedTCM:                                string(runes[12:13]),
		TollExchangeCategory:                       string(runes[13:15]),
		RouteToTelidentMachine:                     string(runes[15:16]),
	}
}
