package cortex

import "strings"

type HeadsetsResponse []HeadsetResponse

type HeadsetResponse struct {
	ID               string   `json:"id"`
	Status           string   `json:"status"`
	ConnectedBy      string   `json:"connectedBy"`
	Dongle           string   `json:"dongle"`
	Firmware         string   `json:"firmware"`
	MotionSensors    []string `json:"motionSensors"`
	Sensors          []string `json:"sensors"`
	Settings         Settings `json:"settings"`
	HeadbandPosition string   `json:"headbandPosition"`
	CustomName       string   `json:"customName"`
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

const (
	noHeadsetConnectedError = "no headset connected"
	isConnectedString       = "has been connected or is connecting"
)

func isConnected(resp ConnectHeadsetResponse) bool {
	if resp.Command != "connect" {
		return false
	}
	if !strings.Contains(resp.Message, isConnectedString) {
		return false
	}
	return true
}

func GetConnectHeadsetRequest(headsetID string) Request {
	return Request{
		ID:      4,
		JsonRPC: "2.0",
		Method:  "controlDevice",
		Params: ControlDeviceParams{
			Command: "connect",
			Headset: headsetID,
		},
	}
}

type ControlDeviceParams struct {
	Command string `json:"command"`
	Headset string `json:"headset"`
}

type ConnectHeadsetResponse struct {
	Command string `json:"command"`
	Message string `json:"message"`
}
