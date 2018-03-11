package main

type acdGroupQue struct {
	DynamicQueue                 string `json:"dynamicQueue"`
	MinimumQueueLength           string `json:"minimumQueueLength"`
	MaximumQueueLength           string `json:"maximumQueueLength"`
	QueueConstant                string `json:"queueConstant"`
	CommonQueueSelectionPriority string `json:"commonQueueSelectionPriority"`
}

func createNewAcdGroupQue(data []rune) acdGroupQue {
	return acdGroupQue{
		DynamicQueue:                 string(data[0:1]),
		MinimumQueueLength:           string(data[1:3]),
		MaximumQueueLength:           string(data[3:6]),
		QueueConstant:                string(data[6:8]),
		CommonQueueSelectionPriority: string(data[8:10]),
	}
}
