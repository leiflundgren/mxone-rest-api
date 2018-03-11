package main

type acdGroupServ struct {
	DirectIndialingTraffic                    string `json:"directIndialingTraffic"`
	CharacteristicsForSwissMFC                string `json:"characteristicsForSwissMFC"`
	MusicOnWait                               string `json:"musicOnWait"`
	CTIGroup                                  string `json:"ctiGroup"`
	CollectCallCategory                       string `json:"collectCallCategory"`
	AutomaticExtendingForExtensionGroupQueues string `json:"automaticExtendingForExtensionGroupQueues"`
}

func createNewAcdGroupServ(data []rune) acdGroupServ {
	return acdGroupServ{
		DirectIndialingTraffic:                    string(data[0:1]),
		CharacteristicsForSwissMFC:                string(data[1:2]),
		MusicOnWait:                               string(data[2:3]),
		CTIGroup:                                  string(data[3:4]),
		CollectCallCategory:                       string(data[4:5]),
		AutomaticExtendingForExtensionGroupQueues: string(data[5:6]),
	}
}
