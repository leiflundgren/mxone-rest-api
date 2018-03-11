package main

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

func CreateNewAcdGroupQue(data string) AcdGroupQue {

	var result AcdGroupQue

	runes := []rune(data)

	result.DynamicQueue = false
	if string(runes[0:0]) == "1" {
		result.DynamicQueue = true
	}

	result.MinimumQueueLength = string(runes[1:4])
	result.MaximumQueueLength = string(runes[4:7])
	result.QueueConstant = string(runes[7:9])
	result.CommonQueueSelectionPriority = string(runes[9:11])

	return result
}
