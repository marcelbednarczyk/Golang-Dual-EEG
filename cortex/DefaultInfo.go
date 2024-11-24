package cortex

type DefaultInfoResponse struct {
	BuildDate              string `json:"buildDate"`
	BuildNumber            string `json:"buildNumber"`
	CloudSyncLogFolderPath string `json:"cloudSyncLogFolderPath"`
	CortexLogFolderPath    string `json:"cortexLogFolderPath"`
	Version                string `json:"version"`
}

func GetDefaultInfoRequest() Request {
	return Request{
		ID:      0,
		JsonRPC: "2.0",
		Method:  "getCortexInfo",
		Params:  nil,
	}
}
