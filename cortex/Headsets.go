package cortex

type HeadsetsResponse []HeadsetResponse

type HeadsetResponse struct {
	ConnectedBy      string        `json:"connectedBy"`
	CustomName       string        `json:"customName"`
	DfuTypes         []interface{} `json:"dfuTypes"`
	Dongle           string        `json:"dongle"`
	Firmware         string        `json:"firmware"`
	FirmwareDisplay  string        `json:"firmwareDisplay"`
	HeadbandPosition interface{}   `json:"headbandPosition"`
	ID               string        `json:"id"`
	IsDfuMode        bool          `json:"isDfuMode"`
	IsVirtual        bool          `json:"isVirtual"`
	MotionSensors    []string      `json:"motionSensors"`
	Sensors          []string      `json:"sensors"`
	Settings         Settings      `json:"settings"`
	Status           string        `json:"status"`
	VirtualHeadsetID string        `json:"virtualHeadsetId"`
}

type Settings struct {
	EegRate  int    `json:"eegRate"`
	EegRes   int    `json:"eegRes"`
	MemsRate int    `json:"memsRate"`
	MemsRes  int    `json:"memsRes"`
	Mode     string `json:"mode"`
}

func GetHeadsetsRequest() Request {
	return Request{
		ID:      2,
		JsonRPC: "2.0",
		Method:  "queryHeadsets",
		Params:  nil,
	}
}
