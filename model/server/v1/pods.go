package v1

type Pod struct {
	Name       string      `json:"podName"`
	Namespace  string      `json:"namespace"`
	Status     Status      `json:"status"`
	Nodename   string      `json:"nodename"`
	Containers []Container `json:"containers"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Status struct {
	Phase     string `json:"phase"`
	HostIP    string `json:"hostIP"`
	PodIP     string `json:"podIP"`
	StartTime string `json:"startTime"`
}
