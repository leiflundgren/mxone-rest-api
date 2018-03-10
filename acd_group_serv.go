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
