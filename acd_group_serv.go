package main

// AcdGroupServ - The parameter states whether direct indialling traffic is permitted/not permitted to the
// group number, whether the A-party will be charged or not on answer and if request of A-number will be
// done or not (only for Swiss MFC DID-trunks), the selection of music-on-wait option, and if the ACD
// group is classified as a CTI group and if the ACD group is permitted to accept collect calls (Brazil)
// and if automatic extending to ACD group queues is permitted or not.
type AcdGroupServ struct {
	DirectIndialingTraffic                                bool `json:"directIndialingTraffic"`
	ASubscriberWillNotBeChargedAndNoANumberRequested      bool `json:"aSubscriberWillNotBeChargedAndNoANumberRequested"`
	ASubscriberWillBeChargedAtAnswer                      bool `json:"aSubscriberWillBeChargedAtAnswer"`
	ANumberInformationIsRequestedToBeDisplayedAtTheBParty bool `json:"aNumberInformationIsRequestedToBeDisplayedAtTheBParty"`
	MusicOnWait                                           bool `json:"musicOnWait"`
	MusicOnWaitOnlyForRecordedAnnouncement                bool `json:"musicOnWaitOnlyForRecordedAnnouncement"`
	MusicOnWaitOnlyForCallsThatReenterQueue               bool `json:"musicOnWaitOnlyForCallsThatReenterQueue"`
	CTIGroupDisplaySelectedMember                         bool `json:"ctiGroupDisplaySelectedMember"`
	CTIGroupDisplayGroupName                              bool `json:"ctiGroupDisplayGroupName"`
	AllowCollectCall                                      bool `json:"allowCollectCall"`
	ExtensionAutomaticExtendingPermittedForQueuedCalls    bool `json:"extensionAutomaticExtendingPermittedForQueuedCalls"`
	ExtensionAutomaticExtendingRequiredForQueuedCalls     bool `json:"extensionAutomaticExtendingRequiredForQueuedCalls"`
}

func CreateNewAcdGroupServ(data string) AcdGroupServ {
	var result AcdGroupServ

	runes := []rune(data)

	if string(runes[0:1]) == "1" {
		result.DirectIndialingTraffic = true
	}

	var didSwissCharacteristics = string(runes[1:2])
	if didSwissCharacteristics == "0" {
		result.ASubscriberWillNotBeChargedAndNoANumberRequested = true
	} else if didSwissCharacteristics == "1" {
		result.ASubscriberWillBeChargedAtAnswer = true
	} else if didSwissCharacteristics == "2" {
		result.ANumberInformationIsRequestedToBeDisplayedAtTheBParty = true
	}

	var musicOnWait = string(runes[2:3])
	if musicOnWait == "1" {
		result.MusicOnWait = true
	} else if musicOnWait == "2" {
		result.MusicOnWaitOnlyForRecordedAnnouncement = true
	} else if musicOnWait == "3" {
		result.MusicOnWaitOnlyForCallsThatReenterQueue = true
	}

	var ctiGroup = string(runes[3:4])
	if ctiGroup == "1" {
		result.CTIGroupDisplaySelectedMember = true
	} else if ctiGroup == "2" {
		result.CTIGroupDisplayGroupName = true
	}

	var collectCallCategory = string(runes[4:5])
	if collectCallCategory == "1" {
		result.AllowCollectCall = true
	}

	var automaticExtendingForExtensionGroupQueues = string(runes[5:6])
	if automaticExtendingForExtensionGroupQueues == "1" {
		result.ExtensionAutomaticExtendingPermittedForQueuedCalls = true
	} else if automaticExtendingForExtensionGroupQueues == "2" {
		result.ExtensionAutomaticExtendingRequiredForQueuedCalls = true
	}

	return result
}
