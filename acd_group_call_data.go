package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
)

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

type AcdGroupSel struct {
	// Member Selection Order. Whether searching for a free member in the group should be undertaken in
	// the order in which the group members were initiated into the ACD group or according to the group
	// member who has been free for the longest time.
	SequentialSearching              bool `json:"sequentialSearching"`
	SequentialSearchingToFreeMembers bool `json:"sequentialSearchingToFreeMembers"`

	// Internal Queuing. Whether internal calls towards the ACD group should be queued or not.
	QueueInternalCallsTowardsTheGroup bool `json:"queueInternalCallsTowardsTheGroup"`

	// Overflow. Whether calls towards the ACD group are permitted or not permitted to be diverted.
	OverflowPermittedNoAvailableMembersOrFull bool `json:"overflowPermittedNoAvailableMembersOrFull"`
	OverflowPermittedWhenFull                 bool `json:"overflowPermittedWhenFull"`
	OverflowNotPermitted                      bool `json:"overflowNotPermitted"`

	// External Overflow/Follow Me. Maximum number of simultaneous calls, which are overflowed to an external
	// destination. Overflow occurs either when there are no queue positions or when all members of the
	// ACD group are unavailable or at both of those conditions.
	MaxNumberOfOverflowsToExternalDestination string `json:"maxNumberOfOverflowsToExternalDestination"`
}

// AcdGroupQue - Queue handling
type AcdGroupQue struct {
	// Whether dynamic queue shall be used or not. At dynamic queue, the queue length will vary depending
	// on number of available group members for the ACD group. It can never be less than the minimum queue length
	// or higher than the maximum queue length.
	DynamicQueue bool `json:"dynamicQueue"`

	// Minimum queue length. Minimum number of delayed calls which can be queued towards an ACD group (used at
	// dynamic queue).
	MinimumQueueLength string `json:"minimumQueueLength"`

	// Maximum queue length. Maximum number of delayed calls which can be queued towards an ACD group.
	MaximumQueueLength string `json:"maximumQueueLength"`

	// Queue constant. The constant is used at dynamic queue to calculate a current queue length, it will
	// alter depending on number of available group members within the ACD group. The value of the constant
	// divided by 10 states the number of queue spaces the current queue length alters.
	QueueConstant string `json:"queueConstant"`

	// Common queue or selection priority. The ACD group's priority. The lower value that is assigned the
	// higher priority.
	CommonQueueSelectionPriority string `json:"commonQueueSelectionPriority"`
}

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

		acdGroupEntry := AcdGroupCallData{
			Group:   grp,
			Lim:     lim,
			servRaw: serv,
			Traf:    traf,
			selRaw:  sel,
			queRaw:  que,
			Cust:    cust,
			Sat:     sat,
			Tpcs:    tpcs,
			Actc:    actc,
		}

		result = append(result, acdGroupEntry)
	}

	return result
}
