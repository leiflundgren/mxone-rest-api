package main

type acdGroupSel struct {
	TypeOfSearching          string `json:"typeOfSearching"`
	HandlingOfInternalCalls  string `json:"handlingOfInternalCalls"`
	OverflowCategory         string `json:"overFlowCategory"`
	ExternalOverflowFollowMe string `json:"externalOverflowFollowMe"`
}

func createNewAcdGroupSel(data []rune) acdGroupSel {
	return acdGroupSel{
		TypeOfSearching:          string(data[0:1]),
		HandlingOfInternalCalls:  string(data[1:2]),
		OverflowCategory:         string(data[2:3]),
		ExternalOverflowFollowMe: string(data[3:5]),
	}
}
