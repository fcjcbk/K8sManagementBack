package v1

type Node struct {
	Name           string `json:"name"`
	CrateTimeStamp string `json:"createTimeStamp"`
	RemainCpu      string `json:"remainCPU"`
	RemainMemory   string `json:"remainMemory"`
}
