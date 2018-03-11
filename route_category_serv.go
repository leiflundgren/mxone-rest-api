package main

type routeCategoryServ struct {
	CallWaitingCharacteristics                               string `json:"callWaitingCharacteristics"`
	AutomaticCallbackCharacteristics                         string `json:"automaticCallbackCharacteristics"`
	TypeOfRoute                                              string `json:"typeOfRoute"`
	CallMeteringCharacteristics                              string `json:"callMeteringCharacteristics"`
	PagingOverSpeechChannel                                  string `json:"pagingOverSpeechChannel"`
	LeastCostRoutingClassOfService                           string `json:"leastCostRoutingClassOfService"`
	MobileExtensionWithoutR1Number                           string `json:"mobileExtensionWithoutR1Number"`
	PresentationOfCallingOrConnectedNumber                   string `json:"presentationOfCallingOrConnectedNumber"`
	RequestCallingNumberFromPSTN                             string `json:"requestCallingNumberFromPSTN"`
	NumberConversionBearerCapabilityAndHighLevelSubstitution string `json:"numberConversionBearerCapabilityAndHighLevelSubstitution"`
}

func createNewRouteCategoryServ(data []rune) routeCategoryServ {
	return routeCategoryServ{
		CallWaitingCharacteristics:                               string(data[0:1]),
		AutomaticCallbackCharacteristics:                         string(data[1:2]),
		TypeOfRoute:                                              string(data[2:3]),
		CallMeteringCharacteristics:                              string(data[3:4]),
		PagingOverSpeechChannel:                                  string(data[4:5]),
		LeastCostRoutingClassOfService:                           string(data[5:6]),
		MobileExtensionWithoutR1Number:                           string(data[6:7]),
		PresentationOfCallingOrConnectedNumber:                   string(data[7:8]),
		RequestCallingNumberFromPSTN:                             string(data[8:9]),
		NumberConversionBearerCapabilityAndHighLevelSubstitution: string(data[9:10]),
	}
}
