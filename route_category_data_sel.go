package main

type directIndialingCharacteristics int

const (
	noReroutingOrNoReroutingAtDIDTraffic  directIndialingCharacteristics = 0
	reroutingOnCongestion                 directIndialingCharacteristics = 1
	reroutingOnBusyOrNotAvailable         directIndialingCharacteristics = 2
	reroutingOnBusyOrNotAvailableIncoming directIndialingCharacteristics = 4
)

func (directIndialing directIndialingCharacteristics) String() string {
	names := [...]string{
		"no-rerouting-or-no-rerouting-at-did-traffic",
		"rerouting-on-congestion",
		"rerouting-on-busy-or-not-available",
		"",
		"rerouting-on-busy-or-not-available-incoming",
	}

	return names[directIndialing]
}

type incomingTraffic int

const (
	barredForIncomingTraffic incomingTraffic = 0
	openForIncomingTraffic   incomingTraffic = 1
)

func (incomingTraffic incomingTraffic) String() string {
	names := [...]string{
		"barred-for-incoming-traffic",
		"open-for-incoming-traffic",
	}

	return names[incomingTraffic]
}

type lineSelectionDuringOutgoingTraffic int

const (
	barredForOutgoingTraffic                 lineSelectionDuringOutgoingTraffic = 0
	evenSeizureInLimOrOriginator             lineSelectionDuringOutgoingTraffic = 1
	sequentialSelectionFirstFreeExternalLine lineSelectionDuringOutgoingTraffic = 2
	sequentialSelectionInRoute               lineSelectionDuringOutgoingTraffic = 3
	itut                                     lineSelectionDuringOutgoingTraffic = 4
)

func (lineSelection lineSelectionDuringOutgoingTraffic) String() string {
	names := [...]string{
		"barred-for-outgoing-traffic",
		"even-seizure-in-lim-or-originator",
		"sequential-selection-first-free-external-line",
		"sequential-selection-in-route",
		"itut",
	}

	return names[lineSelection]
}

type routeCharacteristicsOutgoingTraffic int

const (
	normalRoute   routeCharacteristicsOutgoingTraffic = 0
	terminalRoute routeCharacteristicsOutgoingTraffic = 1
)

func (outgoingCharacteristics routeCharacteristicsOutgoingTraffic) String() string {
	names := [...]string{
		"normal-route",
		"terminal-route",
	}

	return names[outgoingCharacteristics]
}

type alternativeRouteSelectionOnIncomingTraffic int

const (
	alternativeRouteIsPermitted                  alternativeRouteSelectionOnIncomingTraffic = 0
	alternativeRouteIsNotPermitted               alternativeRouteSelectionOnIncomingTraffic = 1
	alternativeRouteIsPermittedWhenRouteIsFaulty alternativeRouteSelectionOnIncomingTraffic = 2
)

type callServiceInformation int

const (
	normalExtension callServiceInformation = 0
	classAExtension callServiceInformation = 1
	classBExtension callServiceInformation = 2
	classCExtension callServiceInformation = 3
	classDExtension callServiceInformation = 4
	dataExtension   callServiceInformation = 5
	pbxOperator     callServiceInformation = 6
	classEextension callServiceInformation = 7
)

func (callInfo callServiceInformation) String() string {
	names := [...]string{
		"normal-extension",
		"class-a-extension",
		"class-b-extension",
		"class-c-extension",
		"class-d-extension",
		"data-extension",
		"pbx-operator",
		"class-e-extension",
	}

	return names[callInfo]
}

type tollExchangeCategory string

const (
	tollExchangeCategoryIsNotUsed                  tollExchangeCategory = "01"
	hospitalityExtensionAuthorized                 tollExchangeCategory = "02"
	extensionAuthorizedForLocalCalls               tollExchangeCategory = "03"
	extensionWithPriorityAuthorized                tollExchangeCategory = "04"
	extensionWithoutPriorityAuthorizedFreeOfCharge tollExchangeCategory = "05"
	trunkCoinBoxOrPostOfficeAuthorized             tollExchangeCategory = "06"
	extensionWithoutPriorityAuthorized             tollExchangeCategory = "07"
	extensionConnectedToDataset                    tollExchangeCategory = "08"
	coinBox                                        tollExchangeCategory = "09"
	dispatcher                                     tollExchangeCategory = "11"
)

func (tollCategory tollExchangeCategory) String() string {
	names := map[tollExchangeCategory]string{
		"01": "toll-exchange-category-is-not-used",
		"02": "hospitality-extension-authorized",
		"03": "extension-authorized-for-local-calls",
		"04": "extension-with-priority-authorized",
		"05": "extension-without-priority-authorized-free-of-charge",
		"06": "trunk-coin-box-or-post-office-authorized",
		"07": "extension-without-priority-authorized",
		"08": "extension-connected-to-dataset",
		"09": "coin-box",
		"11": "dispatcher",
	}

	return names[tollCategory]
}

type routeToTelidentMachine int

const (
	notSetUpForTelident routeToTelidentMachine = 0
	setUpForTelident    routeToTelidentMachine = 1
)

func (telident routeToTelidentMachine) String() string {
	names := [...]string{
		"not-set-up-for-telident-machine",
		"set-up-for-telident",
	}

	return names[telident]
}

type routeCategoryDataSel struct {
	DirectIndialingCharacteristics             string `json:"direct-indialing-characteristics"`
	IncomingTraffic                            string `json:"incoming-traffic"`
	LineSelection                              string `json:"line-selection"`
	RouteCharacteristicsOutgoingTraffic        string `json:"route-characteristics-outgoing-traffic"`
	AlternativeRouteSelectionOnIncomingTraffic string `json:"alternative-route-selection-on-incoming-traffic"`
	CustomerAffiliation                        string `json:"customer-affiliation"`
	VirtualCalls                               bool   `json:"virtual-calls"`
	MaliciousCallTracing                       bool   `json:"malicious-call-tracing"`
	FRLCategoryForIncoming                     int    `json:"frl-category-for-incoming"`
	CallServiceInformation                     string `json:"call-service-information"`
	ReceivedTCM                                bool   `json:"received-tcm"`
	TollExchangeCategory                       string `json:"toll-exchange-category"`
	RouteToTelidentMachine                     string `json:"route-to-telident-machine"`
}

func createNewRouteCategorySel(data string) routeCategoryDataSel {
	runes := []rune(data)

	var result routeCategoryDataSel

	switch string(runes[0:1]) {
	case "0":
		result.DirectIndialingCharacteristics = noReroutingOrNoReroutingAtDIDTraffic.String()
	case "1":
		result.DirectIndialingCharacteristics = reroutingOnCongestion.String()
	case "2":
		result.DirectIndialingCharacteristics = reroutingOnBusyOrNotAvailable.String()
	case "4":
		result.DirectIndialingCharacteristics = reroutingOnBusyOrNotAvailableIncoming.String()
	default:
		// FIXXXXXXXX - What does 7 mean?
		result.DirectIndialingCharacteristics = "not-implemented-yet"
	}

	switch string(runes[1:2]) {
	case "0":
		result.IncomingTraffic = barredForIncomingTraffic.String()
	case "1":
		result.IncomingTraffic = openForIncomingTraffic.String()
	}

	switch string(runes[2:3]) {
	case "0":
		result.LineSelection = barredForOutgoingTraffic.String()
	case "1":
		result.LineSelection = evenSeizureInLimOrOriginator.String()
	case "2":
		result.LineSelection = sequentialSelectionFirstFreeExternalLine.String()
	case "3":
		result.LineSelection = sequentialSelectionInRoute.String()
	case "4":
		result.LineSelection = itut.String()
	}

	switch string(runes[3:4]) {
	case "0":
		result.RouteCharacteristicsOutgoingTraffic = normalRoute.String()
	case "1":
		result.RouteCharacteristicsOutgoingTraffic = terminalRoute.String()
	case "2":
		result.RouteCharacteristicsOutgoingTraffic = "obsolete"
	case "3":
		result.RouteCharacteristicsOutgoingTraffic = "obsolete"
	}

	return result
}
