package main

type routeCategorySel struct {
	DirectIndialingCharacteristics             string `json:"directIndialingCharacteristics"`
	IncomingTraffic                            string `json:"incomingTraffic"`
	LineSelection                              string `json:"lineSelection"`
	RouteCharacteristicsOutgoingTraffic        string `json:"routeCharacteristicsOutgoingTraffic"`
	AlternativeRouteSelectionOnIncomingTraffic string `json:"alternativeRouteSelectionOnIncomingTraffic"`
	CustomerAffiliation                        string `json:"customerAffiliation"`
	VirtualCalls                               string `json:"virtualCalls"`
	MaliciousCallTracing                       string `json:"maliciousCallTracing"`
	FRLCategoryForIncoming                     string `json:"frlCategoryForIncoming"`
	CallServiceInformation                     string `json:"callServiceInformation"`
	ReceivedTCM                                string `json:"receivedTcm"`
	TollExchangeCategory                       string `json:"tollExchangeCategory"`
	RouteToTelidentMachine                     string `json:"routeToTelidentMachine"`
}

func createNewRouteCategorySel(data []rune) routeCategorySel {
	return routeCategorySel{
		DirectIndialingCharacteristics:             string(data[0:1]),
		IncomingTraffic:                            string(data[1:2]),
		LineSelection:                              string(data[2:3]),
		RouteCharacteristicsOutgoingTraffic:        string(data[3:4]),
		AlternativeRouteSelectionOnIncomingTraffic: string(data[4:5]),
		CustomerAffiliation:                        string(data[5:8]),
		VirtualCalls:                               string(data[8:9]),
		MaliciousCallTracing:                       string(data[9:10]),
		FRLCategoryForIncoming:                     string(data[10:11]),
		CallServiceInformation:                     string(data[11:12]),
		ReceivedTCM:                                string(data[12:13]),
		TollExchangeCategory:                       string(data[13:15]),
		RouteToTelidentMachine:                     string(data[15:16]),
	}
}
