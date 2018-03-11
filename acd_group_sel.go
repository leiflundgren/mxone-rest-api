package main

// AcdGroupSel - Selection Category
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
	OverflowNotPermittedWhenFull              bool `json:"overflowPermittedWhenFull"`
	OverflowPermittedNoAvailableMembers       bool `json:"overflowPermittedNoAvailableMembers"`
	OverflowNotPermitted                      bool `json:"overflowNotPermitted"`

	// External Overflow/Follow Me. Maximum number of simultaneous calls, which are overflowed to an external
	// destination. Overflow occurs either when there are no queue positions or when all members of the
	// ACD group are unavailable or at both of those conditions.
	MaxNumberOfOverflowsToExternalDestination string `json:"maxNumberOfOverflowsToExternalDestination"`
}

func CreateNewAcdGroupSel(data string) AcdGroupSel {
	var result AcdGroupSel

	runes := []rune(data)

	var typeOfSearching = string(runes[0:1])
	if typeOfSearching == "0" {
		result.SequentialSearching = true
	} else if typeOfSearching == "1" {
		result.SequentialSearchingToFreeMembers = true
	}

	var handlingOfInternalCalls = string(runes[1:2])
	if handlingOfInternalCalls == "1" {
		result.QueueInternalCallsTowardsTheGroup = true
	}

	var overflowCategory = string(runes[2:3])
	if overflowCategory == "0" {
		result.OverflowPermittedNoAvailableMembersOrFull = true
	} else if overflowCategory == "1" {
		result.OverflowPermittedNoAvailableMembers = true
	} else if overflowCategory == "2" {
		result.OverflowNotPermittedWhenFull = true
	} else if overflowCategory == "3" {
		result.OverflowNotPermitted = true
	}

	return result
}
